# 🖼️ Go Image Downloader

Simple CLI tool written in Go that downloads images concurrently using goroutines and a worker pool.

## 🚀 Features

- Reads image URLs from a text file
- Downloads images concurrently (configurable worker count)
- Saves images to the `images/` directory
- Logs errors (e.g., failed downloads) to `errors.log`
- Uses goroutines, channels, and standard Go libraries

## 📟 Usage

1. Clone the repository:

```bash
git clone https://github.com/Liripol/go-image-downloader.git
cd go-image-downloader
```

2. Add image URLs (one per line) to `urls.txt`

```txt
https://via.placeholder.com/150
https://via.placeholder.com/300
https://via.placeholder.com/600
```

3. Run the program:

```bash
go run main.go
```

4. Downloaded images will appear in the `images/` folder  
   Errors (e.g. network issues) will be written to `errors.log`

## 🗂️ Project Structure

```
go-image-downloader/
├── main.go
├── downloader/
│   └── downloader.go
├── urls.txt
├── images/
├── errors.log
├── .gitignore
└── README.md
```

## 🔧 Tech Stack

- Go (Golang)
- Goroutines & channels
- Standard Go libraries (`net/http`, `io`, `os`, etc.)

## 📌 TODO

- [ ] Add CLI flags (`--file`, `--workers`, etc.)
- [ ] Add download progress tracking
- [ ] Add tests

## 📄 License

📌 Note: This project was built for educational purposes and portfolio demonstration.  
Feel free to explore or reuse parts of the code.
