class Codec:
    def __init__(self) -> None:
        self.id = 0
        self.database = {}

    def encode(self, longUrl:str) -> str:
        self.id += 1
        self.database[self.id] = longUrl
        return 'http://tinyurl.com/'+str(self.id)

    def decode(self, shotUrl:str) -> str:
        i = shotUrl.rfind('/')
        id = int(shotUrl[i+1:])
        return self.database[id]