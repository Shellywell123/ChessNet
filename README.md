# 
```
   █████████  █████                               ██████   █████           █████   
  ███░░░░░███░░███                               ░░██████ ░░███           ░░███    
 ███     ░░░  ░███████    ██████   █████   █████  ░███░███ ░███   ██████  ███████  
░███          ░███░░███  ███░░███ ███░░   ███░░   ░███░░███░███  ███░░███░░░███░   
░███          ░███ ░███ ░███████ ░░█████ ░░█████  ░███ ░░██████ ░███████   ░███    
░░███     ███ ░███ ░███ ░███░░░   ░░░░███ ░░░░███ ░███  ░░█████ ░███░░░    ░███ ███
 ░░█████████  ████ █████░░██████  ██████  ██████  █████  ░░█████░░██████   ░░█████ 
  ░░░░░░░░░  ░░░░ ░░░░░  ░░░░░░  ░░░░░░  ░░░░░░  ░░░░░    ░░░░░  ░░░░░░     ░░░░░  

  █████████████████████████████████████████████████████████████████████████████████
```

A personal project for learning [Go](https://go.dev/) involving [GNUChess](https://www.gnu.org/software/chess/)
# Resources
- [GNUChess Documentation](https://www.gnu.org/software/chess/manual/)
- [Chess Notation](https://www.chess.com/article/view/chess-notation)

# Run in docker contianer
```bash
sh run_in_docker_container.sh
```
In a seperate terminal:
```bash
curl -X POST localhost:3000/game -d '{"move": "e2e4", "user": "tester"}'
```

# Run in local kubernetes cluster
require `kubectl` and `minikube`
```bash
sh run_in_kubernetes_cluster.sh
```
In a seperate terminal:
```bash
curl -X POST <url-shown-in-1st-terminal>/game -d '{"move": "e2e4", "user": "tester"}'
```
