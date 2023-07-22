fn fib(num int) -> float {
    let a = 1.0;
    let b = 0.0;

    for (let i = num; i > 0; i = i - 1) {
        const tmp = a;
        a = a + b;
        b = tmp;
    }

    return b;
}

print(fib(100).to_string() + "\n");
