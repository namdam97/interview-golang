#!/usr/bin/env python3
"""
Script để chạy demo các coding patterns
"""
import subprocess
import sys
import time

def run_command(cmd):
    """Chạy command và hiển thị output"""
    print(f"🚀 Running: {cmd}")
    print("-" * 50)
    
    start_time = time.time()
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
    end_time = time.time()
    
    if result.stdout:
        print(result.stdout)
    if result.stderr:
        print(f"Error: {result.stderr}", file=sys.stderr)
    
    print(f"⏱️  Execution time: {end_time - start_time:.2f}s")
    print("=" * 60)
    return result.returncode

def main():
    print("🎯 GOLANG CODING PATTERNS DEMO")
    print("=" * 60)
    
    commands = [
        "go mod tidy",
        "go run main.go",
        "go run cmd/benchmark/main.go"
    ]
    
    for cmd in commands:
        if run_command(cmd) != 0:
            print(f"❌ Command failed: {cmd}")
            sys.exit(1)
    
    print("✅ All demos completed successfully!")

if __name__ == "__main__":
    main()
