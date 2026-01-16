#!/bin/sh
set -e

# Run pre-runner
./tools/bento-gallery-pre-runner

# Start node application
node build