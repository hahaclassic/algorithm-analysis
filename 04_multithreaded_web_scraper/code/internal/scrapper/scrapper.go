package scrapper

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/hahaclassic/algorithm-analysis/04_multithreaded_web_scrapper/code/pkg/collection"
	"golang.org/x/net/html"
)

var (
	ErrCreateDir   = errors.New("scrapper: can't create directory")
	ErrLoadPage    = errors.New("scrapper: can't load page")
	ErrSavePage    = errors.New("scrapper: can't save page")
	ErrExtractURLs = errors.New("scrapper: can't extract urls")
)

type Scraper struct {
	// countSaved atomic.Int32
	visited   *collection.Collection[string]
	unvisited *collection.Collection[string]
}

func New() *Scraper {
	return &Scraper{
		visited:   collection.New[string](),
		unvisited: collection.New[string](),
	}
}

func (s *Scraper) Start(baseURL, dirPath string, maxWorkers, maxPages int) {
	inputParams(baseURL, dirPath, maxWorkers, maxPages)

	err := createDir(dirPath)
	if err != nil {
		slog.Error("[Start]", "err", err)
		return
	}

	s.start(baseURL, dirPath, maxWorkers, maxPages)

	results(s.visited.Len())
}

func (s *Scraper) start(baseURL, dirPath string, maxWorkers, maxPages int) {
	wg := &sync.WaitGroup{}
	numOfWorkers := 0
	s.unvisited.Store(baseURL)

	for s.visited.Len() < maxPages && s.unvisited.Len() > 0 {
		urls := s.unvisited.Keys()
		s.unvisited.Clear()

		for i := range len(urls) {
			if numOfWorkers >= maxWorkers {
				wg.Wait()
				numOfWorkers = 0
			}

			wg.Add(1)
			numOfWorkers++
			go func() {
				defer wg.Done()
				s.scrapePage(dirPath, maxPages, urls[i])
			}()
		}

		wg.Wait()
	}
}

func (s *Scraper) scrapePage(dirPath string, maxPages int, link string) {
	if s.visited.Load(link) || s.visited.Len() >= maxPages {
		return
	}
	s.visited.Store(link)

	const msg = "[Scrape]"

	body, err := loadPage(link)
	if err != nil {
		slog.Error(msg, "err", err)
		return
	}

	savePath := getSavePath(dirPath, link)
	err = savePage(savePath, strings.NewReader(string(body)))
	if err != nil {
		slog.Error(msg, "err", err)
		return
	}
	// s.countSaved.Add(1)

	err = s.extractURLs(link, strings.NewReader(string(body)))
	if err != nil {
		slog.Error(msg, "err", err)
	}
}

func (s *Scraper) extractURLs(baseLink string, body io.Reader) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w: %w", ErrExtractURLs, err)
		}
	}()

	baseURL, err := url.Parse(baseLink)
	if err != nil {
		return err
	}

	tokenizer := html.NewTokenizer(body)
	tokenType := tokenizer.Next()
	if tokenType == html.ErrorToken {
		return fmt.Errorf("bad body")
	}

	for ; tokenType != html.ErrorToken; tokenType = tokenizer.Next() {
		token := tokenizer.Token()
		if tokenType != html.StartTagToken || token.Data != "a" {
			continue
		}

		for _, attr := range token.Attr {
			if attr.Key != "href" {
				continue
			}

			parsedURL, err := baseURL.Parse(attr.Val)
			if err != nil {
				continue
			}

			if strURL := parsedURL.String(); !s.visited.Load(strURL) {
				s.unvisited.Store(strURL)
			}
		}
	}

	return nil
}

func loadPage(link string) (_ []byte, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("[link=%s] %w: %w", link, ErrLoadPage, err)
		}
	}()

	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func savePage(savePath string, body io.Reader) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w: %w", ErrLoadPage, err)
		}
	}()

	rawBody, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = os.WriteFile(savePath, rawBody, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getSavePath(dirPath string, link string) string {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "invalid_url_" + link + ".html"
	}

	fileName := parsedURL.Path
	if fileName == "" {
		fileName = parsedURL.Host
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9-_]`)
	fileName = re.ReplaceAllString(fileName, "_")
	fileName = strings.Trim(fileName, "_") + ".html"

	if len(fileName) > 255 {
		fileName = fileName[len(fileName)-255:]
	}

	return filepath.Join(dirPath, fileName)
}
