require 'ffi'

module MyMath
  extend FFI::Library

  ffi_lib './libsum.so'

  attach_function :sum, [:int, :int], :int
end

puts MyMath.sum(100, 200)
