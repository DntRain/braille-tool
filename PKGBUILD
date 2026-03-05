# 维护者: DontRain <xykk0257@gmail.com>
# 这是一个为 Braille-Art 转换器编写的 PKGBUILD

pkgname=braille-tool-git
_pkgname=braille-tool
pkgver=0.1.0
pkgrel=1
pkgdesc="一个用 Go 编写的将图片转换为盲文点阵艺术的命令行工具"
arch=('x86_64')
url="https://github.com/DontRain/${_pkgname}" # 替换为你的 GitHub 仓库地址
license=('MIT')
depends=('glibc')
makedepends=('go')
provides=("${_pkgname}")
conflicts=("${_pkgname}")
# 因为是 -git 包，我们通常从源码构建
source=("${_pkgname}::git+${url}.git")
sha256sums=('SKIP')

build() {
  cd "${_pkgname}"
  
  # 设置 Go 环境变量进行优化编译
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"

  go build -o "${_pkgname}" main.go
}

package() {
  cd "${_pkgname}"
  
  # 将二进制文件安装到系统路径 /usr/bin/
  install -Dm755 "${_pkgname}" "${pkgdir}/usr/bin/${_pkgname}"
  
  # (可选) 如果你有 LICENSE 文件，也一并安装
  # install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
