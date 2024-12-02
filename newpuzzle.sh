
YEAR=$1
DAY=$2

mkdir -p "${YEAR}/${DAY}"

cat > main.go <<- EOM
package main

import(
	"github.com/tmeire/adventofcode/${YEAR}/${DAY}"
)

func main() {
	${DAY}.Solve()
}
EOM

cat > "${YEAR}/${DAY}/${DAY}a.go" <<- EOM
package $DAY

func Solve() {
}
EOM

gofmt -w main.go
gofmt -w "${YEAR}/${DAY}/${DAY}a.go"