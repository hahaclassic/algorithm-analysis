.PHONY: test report_prev
.DEFAULT: unit

TEST_DIR := test

unit: 
	python3 ./unit.py

report-prev: $(TEST_DIR)
	python3 ./unit.py --output=test/stud-unit-test-report-prev.json

$(TEST_DIR):
	@mkdir -p ./test