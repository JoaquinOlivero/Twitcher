echo System updating...
sudo apt update && sudo apt upgrade

echo Getting all the dependencies...
sudo apt install -y autoconf \
automake \
build-essential \
cmake \
git-core \
libfreetype6-dev \
libgnutls28-dev \
libmp3lame-dev \
libtool \
libvorbis-dev \
meson \
ninja-build \
pkg-config \
texinfo \
wget \
yasm \
zlib1g-dev \
libzmq3-dev \

sudo apt install libunistring-dev
sudo apt install nasm

echo Getting libx264...
sudo apt install libx264-dev

echo Getting libvpx...
sudo apt install libvpx-dev

echo Getting aac...
sudo apt install libfdk-aac-dev

echo Getting libopus...
sudo apt install libopus-dev

echo making the necessary folders...
mkdir -pv ~/ffmpeg_sources ~/ffmpeg_build

git clone https://git.ffmpeg.org/ffmpeg.git ffmpeg
cd ffmpeg && \
PATH="$HOME/bin:$PATH" PKG_CONFIG_PATH="$HOME/ffmpeg_build/lib/pkgconfig" ./configure \
  --disable-optimizations \
  --prefix="$HOME/ffmpeg_build" \
  --extra-cflags="-I$HOME/ffmpeg_build/include" \
  --extra-ldflags="-L$HOME/ffmpeg_build/lib" \
  --extra-libs="-lpthread -lm" \
  --ld="g++" \
  --bindir="$HOME/bin" \
  --disable-debug \
  --disable-ffplay \
  --disable-indev=sndio \
  --disable-outdev=sndio \
  --disable-doc \
  --disable-htmlpages \
  --disable-manpages \
  --disable-podpages \
  --disable-txtpages \
  --enable-gpl \
  --enable-gnutls \
  --enable-libfdk-aac \
  --enable-libfreetype \
  --enable-libmp3lame \
  --enable-libopus \
  --enable-libvorbis \
  --enable-libvpx \
  --enable-libx264 \
  --enable-libzmq \
  --enable-nonfree && \
PATH="$HOME/bin:$PATH" make -j2
# make install && \
# hash -r