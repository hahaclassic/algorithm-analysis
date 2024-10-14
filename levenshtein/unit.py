import unittest
import coverage
import json
from datetime import datetime
import pytz
import argparse
import typing
import io
import sys

SRC_DIR = ["./src"]
TEST_DIR = "./"
LOG_FILE = "log.txt"

def RunUnitTests() -> dict[str, typing.Any]:
    log_stream = io.StringIO()
    log_file = open(LOG_FILE, 'w')

    cov = coverage.Coverage(source=SRC_DIR)
    cov.start()
    loader = unittest.TestLoader()
    suite = loader.discover(TEST_DIR)  
    runner = unittest.TextTestRunner(stream=log_stream, verbosity=2)
    result = runner.run(suite)
    cov.stop()
    cov.save()

    log_file.write(log_stream.getvalue())

    failed = len(result.failures) + len(result.errors)
    passed = result.testsRun - failed
    coverage_percent = round(cov.report(file=log_file, show_missing=False))
    tz = pytz.timezone('Europe/Moscow')
    timestamp = datetime.now(tz).strftime("%Y-%m-%dT%H:%M:%S:%z")

    log_stream.close()
    log_file.close()

    report = {
        "timestamp": timestamp,
        "coverage": coverage_percent,
        "passed": passed,
        "failed": failed
    }

    return report

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--output",
        type=str,
        help="the path to the file for saving the unit test report",
        default=""
    )

    report = RunUnitTests()

    args = parser.parse_args()

    if args.output == "":
        json.dump(report, sys.stdout, indent=4)
        print()
    else:
        with open(args.output, 'w') as f:
            json.dump(report, f, indent=4)

