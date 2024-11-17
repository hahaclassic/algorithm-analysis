READY_DIR := ready

.PHONY: clean $(READY_DIR)/app-cli-debug

$(READY_DIR)/report.pdf: $(READY_DIR)
	cp report/report.pdf $(READY_DIR)/report.pdf

$(READY_DIR)/stud-unit-test-report-prev.json: $(READY_DIR)
	cp test/stud-unit-test-report-prev.json $(READY_DIR)/stud-unit-test-report-prev.json

$(READY_DIR)/stud-unit-test-report.json: $(READY_DIR)
	cp test/stud-unit-test-report-prev.json $(READY_DIR)/stud-unit-test-report.json

$(READY_DIR)/main-cli-debug.py: $(READY_DIR)

$(READY_DIR)/app-cli-debug: $(READY_DIR)
	go -C ./code build -o ../$(READY_DIR)/app-cli-debug ./cmd/main.go

$(READY_DIR):
	@mkdir -p ./ready

clean:
	rm -rf log.txt ready/ data/
