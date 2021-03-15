# 1-init-vs-main
The Go language reserve two functions for special purpose and the functions are `main()` and `init()` function.

## main() function
- The entry point of the executable programs. 
- It does not take any argument nor return anything. 
- Go automatically call main() function, so there is no need to call main() function explicitly 
- Every executable program must contain single main package and main() function.

## init() function 
- Like main() function, it does not take any argument nor return anything. 
- Called when the package is initialized. 
- This function is declared implicitly, so you cannot reference it from anywhere 
- Allowed to create multiple init() function in the same program and they execute in the order they are created. 
- init() function is executed before the main() function call
- The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context.

<img alt="init-main-flow" src="https://github.com/patrickalexander-tiket/introduction-to-go/blob/master/readme-files/init-main-flow.png" width="auto"> <br/>
