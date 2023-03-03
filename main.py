import sys, subprocess
from config import boot_logo

def GnuChessWrapper():
    """
    Wrapper to run commands in GnuChess CLI
    """
    p = subprocess.Popen('gnuchess', stdin=subprocess.PIPE, stdout=sys.stdout)
    for x in ('e4 ', 'pgnsave game.pgn', 'quit'):
        # p.wait()
        p.stdin.write( x.encode() + b'\n' )
        
    p.communicate()
    
    # currently crashes after waiting for gnuchess to make its move (repeately enters e4?)


def main():
    """
    Main Function (entry point)
    """
    print(boot_logo)
    GnuChessWrapper()   
    

if __name__ == "__main__":
    main()
    