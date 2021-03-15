# 2-private-vs-public

When importing a package you can access only its exported identifiers.
An identifier is exported if it begins with a capital letter.

```
package vehicle

// This will be exported. It is essentially public to other packages that import package shape
// because its first letter is lowercase means public function
type Car struct {}

// This will not be exported. It is essentially private and inaccessible to packages that import package shape
// because its first letter is lowercase means private function
type truck struct {}
```