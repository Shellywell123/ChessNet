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

# start the server
```bash
sh run.sh
```

# request the api
```
curl -X POST localhost:8090/chessnet -d '{"move": "e2e4", "user": "tester"}'
```
# kubernetes
```bash
minikube start
```
```bash
docker build . -o ChessNet.exe
```
```bash
eval $(minikube -p minikube docker-env)
```
```bash
kuberctl apply -f kubernetes/development.yaml
```