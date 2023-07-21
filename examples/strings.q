const str = "testing123!";

for (let i = 0; i < str.len(); i = i + 1) {
    print(str.at(i).to_string());
    print("\n");
}

print(str.slice(2, 5) + "\n");
print(str.replace("123", "456") + "\n");
