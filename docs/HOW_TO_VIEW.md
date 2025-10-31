# 📖 Cara Melihat Dokumentasi

Ada beberapa cara untuk membaca dan melihat preview dokumentasi ini:

---

## 🌐 Opsi 1: Buka di Browser (Paling Mudah)

### **File HTML (Recommended untuk Preview)**

Buka file ini di browser untuk tampilan yang lebih cantik:

```
docs/index.html
```

**Cara:**
1. Klik kanan pada file `index.html`
2. Pilih "Open with" → Browser favorit (Chrome, Firefox, Safari, Edge)
3. Atau double-click file `index.html`

**Tampilan:**
- ✅ Desain yang menarik
- ✅ Navigasi mudah
- ✅ Link ke semua dokumentasi
- ✅ Table of contents
- ✅ Color-coded sections

---

## 📝 Opsi 2: Preview Markdown di VSCode

### **Untuk Developer yang menggunakan VSCode:**

1. **Install Extension (jika belum):**
   - Buka VSCode
   - Extensions (Ctrl+Shift+X / Cmd+Shift+X)
   - Install "Markdown Preview Enhanced" atau "Markdown All in One"

2. **Preview File Markdown:**
   - Buka file `.md` (misalnya `INDEX.md`)
   - Tekan `Ctrl+Shift+V` (Windows/Linux) atau `Cmd+Shift+V` (Mac)
   - Atau klik kanan → "Open Preview"

3. **Split View (Side by Side):**
   - Tekan `Ctrl+K V` (Windows/Linux) atau `Cmd+K V` (Mac)
   - Edit di kiri, preview di kanan

---

## 📱 Opsi 3: Baca di GitHub/GitLab

Jika project sudah di-upload ke repository:

1. Navigate ke folder `docs/` di GitHub/GitLab
2. Klik file `.md` yang ingin dibaca
3. GitHub/GitLab akan otomatis render markdown dengan baik

---

## 🖥️ Opsi 4: Command Line dengan Tools

### **grip (GitHub Readme Instant Preview)**

```bash
# Install
pip install grip

# Run (dari folder docs)
cd docs
grip INDEX.md

# Buka browser: http://localhost:6419
```

### **markdown-preview**

```bash
# Install
npm install -g markdown-preview

# Run
markdown-preview INDEX.md
```

### **mdv (Markdown Viewer - Terminal)**

```bash
# Install
pip install mdv

# View in terminal
mdv INDEX.md
```

---

## 📂 Struktur File Dokumentasi

Semua dokumentasi ada di folder `docs/`:

```
docs/
├── index.html              ← Buka ini di browser! (HTML Preview)
├── 00_START_HERE.md        ← Panduan navigasi (Markdown)
├── INDEX.md                ← Navigation guide lengkap
├── QUICKSTART.md           ← Setup 5 menit
├── README.md               ← Dokumentasi utama
├── API_TESTING.md          ← Contoh testing API
├── ARCHITECTURE_DIAGRAM.md ← Visual diagrams
├── PROJECT_STRUCTURE.md    ← Struktur project
├── DEPLOYMENT.md           ← Panduan deployment
├── IMPLEMENTATION_SUMMARY.md ← Detail implementasi
└── PROJECT_MANIFEST.txt    ← Checklist files
```

---

## 🎯 Recommended Reading Order

### **1. Pertama Kali? Mulai Dari:**
```
docs/index.html (buka di browser)
atau
docs/00_START_HERE.md (baca di VSCode/editor)
```

### **2. Ingin Setup Cepat?**
```
docs/QUICKSTART.md
```

### **3. Ingin Memahami Arsitektur?**
```
docs/ARCHITECTURE_DIAGRAM.md
docs/PROJECT_STRUCTURE.md
```

### **4. Ingin Test API?**
```
docs/API_TESTING.md
```

---

## 💡 Tips untuk Preview Terbaik

### **Di Browser (HTML)**
- ✅ Tampilan paling bagus
- ✅ Navigasi mudah dengan click
- ✅ Responsive design
- ⚠️ File: `docs/index.html`

### **Di VSCode (Markdown)**
- ✅ Edit dan preview side-by-side
- ✅ Link antar dokumen bekerja
- ✅ Good untuk development
- ⚠️ Perlu extension untuk preview yang bagus

### **Di GitHub (Online)**
- ✅ Render otomatis
- ✅ Shareable link
- ✅ Code highlighting
- ⚠️ Perlu upload ke repo dulu

---

## 🔗 Link Navigation

### **Dari Root Project:**
```
📁 /
├── README.md          → Overview project
└── docs/
    ├── index.html     → Preview portal
    └── 00_START_HERE.md → Panduan navigasi
```

### **Internal Links di Dokumentasi:**
Semua file `.md` menggunakan relative links, jadi klik link akan membuka file terkait.

Contoh:
- `[QUICKSTART](QUICKSTART.md)` → Buka file QUICKSTART.md
- `[API Testing](API_TESTING.md)` → Buka file API_TESTING.md

---

## 🎨 Keunggulan Setiap Format

### **HTML (index.html)**
```
✅ Visual yang menarik
✅ Color-coded sections
✅ Interactive navigation
✅ Table of contents
✅ Responsive design
✅ Professional look
```

### **Markdown (.md files)**
```
✅ Easy to edit
✅ Git-friendly (diff tracking)
✅ Portable
✅ Universal format
✅ Works everywhere
✅ GitHub/GitLab compatible
```

---

## 📞 Quick Access

### **Ingin langsung praktek?**
```bash
# Baca ini:
docs/QUICKSTART.md

# Lalu jalankan:
go run ./cmd/server
```

### **Ingin lihat semua endpoint?**
```bash
# Baca ini:
docs/API_TESTING.md
```

### **Ingin deploy?**
```bash
# Baca ini:
docs/DEPLOYMENT.md
```

---

## 🚀 Recommended Tools

### **Untuk Viewing:**
1. **Browser** (Chrome, Firefox, Safari) - Untuk `index.html`
2. **VSCode** + Markdown Extensions - Untuk `.md` files
3. **grip** - Untuk GitHub-style rendering

### **Untuk Editing:**
1. **VSCode** - Best for markdown editing
2. **Typora** - WYSIWYG markdown editor
3. **Mark Text** - Open-source markdown editor

---

## 📋 Checklist: Sudah Lihat Semua?

Tandai yang sudah dibaca:

- [ ] `index.html` - HTML preview portal
- [ ] `00_START_HERE.md` - Panduan navigasi
- [ ] `QUICKSTART.md` - Setup guide
- [ ] `README.md` - Main documentation
- [ ] `API_TESTING.md` - API examples
- [ ] `ARCHITECTURE_DIAGRAM.md` - Visual diagrams
- [ ] `PROJECT_STRUCTURE.md` - Code organization
- [ ] `DEPLOYMENT.md` - Deployment guide
- [ ] `IMPLEMENTATION_SUMMARY.md` - Implementation details
- [ ] `PROJECT_MANIFEST.txt` - File checklist

---

## 🎓 Best Practice

### **Untuk Membaca:**
1. Buka `index.html` di browser untuk overview
2. Klik link ke dokumen yang dibutuhkan
3. Atau buka langsung file `.md` di VSCode

### **Untuk Development:**
1. Keep VSCode open dengan preview
2. Edit code sambil baca dokumentasi
3. Test sambil lihat `API_TESTING.md`

### **Untuk Sharing:**
1. Share link `docs/index.html` untuk quick preview
2. Share specific `.md` files untuk detail
3. Upload ke GitHub untuk public access

---

**Happy Reading! 📚**

**[← Back to Documentation Index](INDEX.md)**
**[🏠 Back to Project Root](../README.md)**
