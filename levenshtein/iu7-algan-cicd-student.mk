READY_DIR := ready

$(READY_DIR)/report.pdf: $(READY_DIR)
	cp doc/report.pdf $(READY_DIR)/report.pdf

$(READY_DIR)/stud-unit-test-report-prev.json: $(READY_DIR)
	cp test/stud-unit-test-report-prev.json $(READY_DIR)/stud-unit-test-report-prev.json

$(READY_DIR)/main-cli-debug.py: $(READY_DIR)
	cp -rf src/* $(READY_DIR)/

$(READY_DIR)/stud-unit-test-report.json: $(READY_DIR)
	python3 ./unit.py --output=$(READY_DIR)/stud-unit-test-report.json

$(READY_DIR):
	@mkdir -p ./ready

.PHONY: clean
clean:
	rm -f .coverage log.txt
