void uint32_to_ordered_bytes(unsigned int u32, char* bytes) {
  for (int i = 0; i < 4; i++) {
    bytes[3 - i] = u32 & 0xff;
    u32 >>= 8;
  }
}

unsigned int ordered_bytes_to_uint32(const char* bytes) {
  unsigned int u32 = 0;
  for (int i = 0; i < 4; i++) {
    u32 <<= 8;
    u32 |= (bytes[i] & 0xff);
  }
  return u32;
}

void int32_to_ordered_bytes(int i32, char* bytes) {
  uint32_to_ordered_bytes(i32 ^ 0x80000000, bytes);
}

int ordered_bytes_to_int32(const char* bytes) {
  return ordered_bytes_to_uint32(bytes) ^ 0x80000000;
}

void uint64_to_ordered_bytes(unsigned long long u64, char* bytes) {
  for (int i = 0; i < 8; i++) {
    bytes[7 - i] = u64 & 0xff;
    u64 >>= 8;
  }
}

unsigned long long ordered_bytes_to_uint64(const char* bytes) {
  unsigned long long u64 = 0;
  for (int i = 0; i < 8; i++) {
    u64 <<= 8;
    u64 |= (bytes[i] & 0xff);
  }
  return u64;
}

void int64_to_ordered_bytes(long long i64, char* bytes) {
  uint64_to_ordered_bytes(i64 ^ 0x8000000000000000ULL, bytes);
}

long long ordered_bytes_to_int64(const char* bytes) {
  return ordered_bytes_to_uint64(bytes) ^ 0x8000000000000000ULL;
}

void float32_to_ordered_bytes(float f32, char* bytes) {
  unsigned int u32 = *(unsigned int*)&f32;
  if ((u32 & 0x80000000) != 0) {
    u32 = ~u32 + 1;
  } else {
    u32 |= 0x80000000;
  }
  uint32_to_ordered_bytes(u32, bytes);
}

float ordered_bytes_to_float32(const char* bytes) {
  unsigned int u32 = ordered_bytes_to_uint32(bytes);
  if ((u32 & 0x80000000) != 0) {
    u32 ^= 0x80000000;
  } else {
    u32 = ~u32 + 1;
  }
  return *(float*)&u32;
}

void float64_to_ordered_bytes(double f64, char* bytes) {
  unsigned long long u64 = *(unsigned long long*)&f64;
  if ((u64 & 0x8000000000000000ULL) != 0) {
    u64 = ~u64 + 1;
  } else {
    u64 |= 0x8000000000000000ULL;
  }
  uint64_to_ordered_bytes(u64, bytes);
}

double ordered_bytes_to_float64(const char* bytes) {
  unsigned long long u64 = ordered_bytes_to_uint64(bytes);
  if ((u64 & 0x8000000000000000ULL) != 0) {
    u64 ^= 0x8000000000000000ULL;
  } else {
    u64 = ~u64 + 1;
  }
  return *(double*)&u64;
}
