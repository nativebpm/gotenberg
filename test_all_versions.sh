#!/bin/bash

# List of Gotenberg versions to test
VERSIONS=(
    "8.23.2" # baseline
    "8.24.0"
    "8.25.0"
    "8.26.0"
    "8.27.0"
    "8.28.0"
    "8.29.0"
    "8.29.1"
    "8.30.0"
    "8.30.1"
    "8.31.0"
    "8.32.0"
    "8.33.0"
)

echo "=== Gotenberg Client Version Test Suite ==="

# Clean up any existing gotenberg container
docker rm -f gotenberg &>/dev/null

for VERSION in "${VERSIONS[@]}"; do
    echo ""
    echo "--------------------------------------------------"
    echo "Testing Gotenberg version: $VERSION"
    echo "--------------------------------------------------"

    # Run the docker container
    docker run -d --name gotenberg -p 3000:3000 gotenberg/gotenberg:$VERSION &>/dev/null
    
    if [ $? -ne 0 ]; then
        echo "Failed to start Docker container for version $VERSION"
        continue
    fi

    # Wait for the service to be healthy (up to 15 seconds)
    HEALTHY=false
    for i in {1..15}; do
        if curl -s http://localhost:3000/health | grep -q '"status":"up"'; then
            HEALTHY=true
            break
        fi
        sleep 1
    done

    if [ "$HEALTHY" = false ]; then
        echo "Gotenberg service is not healthy or did not start in time."
        docker logs gotenberg
        docker rm -f gotenberg &>/dev/null
        continue
    fi

    echo "Gotenberg service is up and running."

    # Clean up old output files
    rm -f out.pdf converturl-chromium.pdf output.pdf merged.pdf

    # Run examples and track success
    FAILED=0

    echo "Running examples/cmd/health..."
    go run examples/cmd/health/main.go > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo "  [FAIL] health example"
        FAILED=$((FAILED + 1))
    else
        echo "  [OK] health example"
    fi

    echo "Running examples/cmd/chromium/helloworld..."
    go run examples/cmd/chromium/helloworld/main.go > /dev/null 2>&1
    if [ $? -ne 0 ] || [ ! -f out.pdf ]; then
        echo "  [FAIL] helloworld example"
        FAILED=$((FAILED + 1))
    else
        echo "  [OK] helloworld example"
    fi

    echo "Running examples/cmd/chromium/converturl..."
    go run examples/cmd/chromium/converturl/main.go > /dev/null 2>&1
    if [ $? -ne 0 ] || [ ! -f converturl-chromium.pdf ]; then
        echo "  [FAIL] converturl example"
        FAILED=$((FAILED + 1))
    else
        echo "  [OK] converturl example"
    fi

    echo "Running examples/cmd/libreoffice/convert..."
    go run examples/cmd/libreoffice/convert/main.go > /dev/null 2>&1
    if [ $? -ne 0 ] || [ ! -f output.pdf ]; then
        echo "  [FAIL] libreoffice convert example"
        FAILED=$((FAILED + 1))
    else
        echo "  [OK] libreoffice convert example"
    fi

    echo "Running examples/cmd/pdfengines/merge..."
    go run examples/cmd/pdfengines/merge/main.go > /dev/null 2>&1
    if [ $? -ne 0 ] || [ ! -f merged.pdf ]; then
        echo "  [FAIL] pdfengines merge example"
        FAILED=$((FAILED + 1))
    else
        echo "  [OK] pdfengines merge example"
    fi

    # Clean up output files
    rm -f out.pdf converturl-chromium.pdf output.pdf merged.pdf

    # Report results for this version
    if [ $FAILED -eq 0 ]; then
        echo "RESULT: Version $VERSION passed all tests successfully!"
    else
        echo "RESULT: Version $VERSION FAILED $FAILED tests."
    fi

    # Stop and remove container
    docker rm -f gotenberg &>/dev/null
done

echo ""
echo "=== Test Suite Completed ==="
