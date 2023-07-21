let echo = "hello world";

fn test(a int, b int) -> int {
    print echo + "\n";
    return a + b;
}

// echo = echo.len();

print test(1, 56);
print "\n";
print echo.len();
print "\n";

if (echo.len() == 12) {
    print "hi!";
}