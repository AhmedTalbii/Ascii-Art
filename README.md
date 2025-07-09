# ASCII Art Generator

[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org)  
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](./LICENSE)  
[![Build Status](https://img.shields.io/github/actions/workflow/status/AhmedTalbii/Ascii-Art/go.yml)](https://github.com/AhmedTalbii/Ascii-Art/actions)

A simple command-line tool that converts your text into stylish ASCII art using a predefined font.

---

## 📸 Preview

![ASCII Art Example](./assets/asciiArt.png)

---

## 🚀 Features

- Converts any input string into a multi-line ASCII art banner  
- Respects `\n` for multi-line input  
- Zero dependencies beyond Go’s standard library  
- Cross-platform: Linux, macOS, Windows

---

## 💾 Installation

1. Make sure [Go](https://golang.org/dl/) (≥1.20) is installed  
2. Clone and build:

   ```bash
   git clone https://github.com/AhmedTalbii/Ascii-Art.git
   cd Ascii-Art
   go build -o ascii-art-generator main.go
   ```

3. (Optional) Move the binary into your `PATH`:

   ```bash
   sudo mv ascii-art-generator /usr/local/bin/
   ```

---

## 🎯 Usage

```bash
# Single-line text
ascii-art-generator "Hello, World!"

# Multi-line support
ascii-art-generator "Line one\nLine two\nLine three"
```

**Example Output:**

```bash
$ ascii-art-generator "GoLang"
  _____       _             
 / ____|     | |            
| |  __  ___ | |_ ___  ___  
| | |_ |/ _ \| __/ _ \/ __| 
| |__| | (_) | ||  __/\__ \ 
 \_____|\___/ \__\___||___/ 
```

---

## ⚙️ Configuration & Custom Fonts

> _Coming soon:_ Ascii Art web is a web site that generates Ascii Art

## You can Do what you ant with the code Enjoy 😃
