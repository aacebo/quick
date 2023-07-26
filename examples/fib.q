fn fib(num float) -> float {
    if (num <= 2.0) {
        return 1.0;
    }

    return fib(num - 1.0) + fib(num - 2.0);
}

print(fib(30.0).to_string() + "\n");
