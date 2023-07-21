for (let i = 0; i < 10; i = i + 1) {
    print i;
    print " -> ";
    print "\"hi\"";
    print "\n";
}

const str = "testing123!";

for (let i = 0; i < str.len(); i = i + 1) {
    print str.at(i);
    print "\n";
}

print str.slice(2, 5) + "\n";