package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// these are the initial values of the buffers
var initial_hash_values []string = []string{
	"6a09e667", "bb67ae85", "3c6ef372", "a54ff53a",
	"510e527f", "9b05688c", "1f83d9ab", "5be0cd19",
}

// these are the constants
var sha_256_constants []string = []string{
	"428a2f98", "71374491", "b5c0fbcf", "e9b5dba5",
	"3956c25b", "59f111f1", "923f82a4", "ab1c5ed5",
	"d807aa98", "12835b01", "243185be", "550c7dc3",
	"72be5d74", "80deb1fe", "9bdc06a7", "c19bf174",
	"e49b69c1", "efbe4786", "0fc19dc6", "240ca1cc",
	"2de92c6f", "4a7484aa", "5cb0a9dc", "76f988da",
	"983e5152", "a831c66d", "b00327c8", "bf597fc7",
	"c6e00bf3", "d5a79147", "06ca6351", "14292967",
	"27b70a85", "2e1b2138", "4d2c6dfc", "53380d13",
	"650a7354", "766a0abb", "81c2c92e", "92722c85",
	"a2bfe8a1", "a81a664b", "c24b8b70", "c76c51a3",
	"d192e819", "d6990624", "f40e3585", "106aa070",
	"19a4c116", "1e376c08", "2748774c", "34b0bcb5",
	"391c0cb3", "4ed8aa4a", "5b9cca4f", "682e6ff3",
	"748f82ee", "78a5636f", "84c87814", "8cc70208",
	"90befffa", "a4506ceb", "bef9a3f7", "c67178f2",
}

// this function will take the decimal number and will convert binary with required number of bits length
func format(dec int64, bits int) string {
	b := strconv.FormatInt(dec, 2)
	z := bits - len(b)
	for i := 0; i < z; i++ {
		b = "0" + b
	}
	return b
}

// this function is to convert the decimal number into 8 bit binary string
func bin_8bit(dec int64) string {
	return format(dec, 8)
}

// this function is to convert the decimal numner into 32 bit binary string
func bin_32bit(dec int64) string {
	return format(dec, 32)
}

// this funcion is to convert the decimal number into 64 bit binary
func bin_64bit(dec int64) string {
	return format(dec, 64)
}

// this function is to convert the hexadecimal string into decimal format
func dec_bin_hex(hex_string string) int64 {
	s, _ := strconv.ParseInt(hex_string, 16, 64)
	return s
}

// this function will take decimal number and returns the hexadecimal value
func hex_return(dec int64) string {
	s := strconv.FormatInt(dec, 16)
	return s
}

// this function is to combine the list of binary string into one single string
func l_s(bit_list []string) string {
	bit_string := ""
	for _, val := range bit_list {
		bit_string = bit_string + val
	}
	return bit_string
}

// this function will take the input string a convert each
// character into ascii value and then convert that to binary
// this will return a binary string
func message_bit_return(string_input string) string {
	bit_list := []string{}
	for i := 0; i < len(string_input); i++ {
		bit_list = append(bit_list, bin_8bit(int64(string_input[i])))
	}
	return l_s(bit_list)
}

// this will add 1 bit to the original message and then add zeros and then
// at last adds the length of the message as binary bits
func message_pad(bit_list string) string {
	pad_one := bit_list + "1"
	pad_len := len(pad_one)
	k := 0
	for ((pad_len+k)-448)%512 != 0 {
		k += 1
	}
	back_append_0 := ""
	for i := 0; i < k; i++ {
		back_append_0 += "0"
	}
	back_append_1 := bin_64bit(int64(len(bit_list)))
	return pad_one + back_append_0 + back_append_1
}

// returns the padded message
func message_pre_pro(input_string string) string {
	bit_main := message_bit_return(input_string)
	return message_pad(bit_main)
}

// to split the padded message into 32bit each
func L_P(SET string, n int) []string {
	to_return := []string{}
	j := 0
	k := n
	for k < len(SET)+1 {
		to_return = append(to_return, SET[j:k])
		j = k
		k += n
	}
	return to_return
}

// this function will returns the list of binary string with each element as 32 bit size
func message_parsing(input_string string) []string {
	return L_P(message_pre_pro(input_string), 32)
}

// this function takes 2 string a will give the XOR of 2 strings
func xor_2str(bit_string_1, bit_string_2 string) string {
	xor_list := []string{}
	for i := 0; i < len(bit_string_1); i++ {
		if bit_string_1[i] == '0' && bit_string_2[i] == '0' {
			xor_list = append(xor_list, "0")
		}
		if bit_string_1[i] == '1' && bit_string_2[i] == '1' {
			xor_list = append(xor_list, "0")
		}
		if bit_string_1[i] == '0' && bit_string_2[i] == '1' {
			xor_list = append(xor_list, "1")
		}
		if bit_string_1[i] == '1' && bit_string_2[i] == '0' {
			xor_list = append(xor_list, "1")
		}
	}
	return l_s(xor_list)
}

// this function takes to strings and performs AND operation
func and_2str(bit_string_1, bit_string_2 string) string {
	and_list := []string{}
	for i := 0; i < len(bit_string_1); i++ {
		if bit_string_1[i] == '1' && bit_string_2[i] == '1' {
			and_list = append(and_list, "1")
		} else {
			and_list = append(and_list, "0")
		}
	}
	return l_s(and_list)
}

// this function takes a string and performs NOT operation on that
func not_str(bit_string string) string {
	not_list := []string{}
	for i := 0; i < len(bit_string); i++ {
		if bit_string[i] == '0' {
			not_list = append(not_list, "1")
		} else {
			not_list = append(not_list, "0")
		}
	}
	return l_s(not_list)
}

// this function will take a string and
func s_l(bit_string string) []string {
	bit_list := []string{}
	for i := 0; i < len(bit_string); i++ {
		bit_list = append(bit_list, string(bit_string[i]))
	}
	return bit_list
}

// this function will right rotate the bits of the string by 'n' bits
func rotate_right(bit_string string, n int) string {
	bit_list := s_l(bit_string)
	length := len(bit_list)
	for i := 0; i < n; i++ {
		last := bit_list[length-1]
		for j := length - 1; j > 0; j-- {
			bit_list[j] = bit_list[j-1]
		}
		bit_list[0] = last
	}
	return l_s(bit_list)
}
func e_0(x string) string {
	return xor_2str(xor_2str(rotate_right(x, 2), rotate_right(x, 13)), rotate_right(x, 22))
}
func e_1(x string) string {
	return xor_2str(xor_2str(rotate_right(x, 6), rotate_right(x, 11)), rotate_right(x, 25))
}
func Ch(x, y, z string) string {
	return xor_2str(and_2str(x, y), and_2str(not_str(x), z))
}
func Maj(x, y, z string) string {
	return xor_2str(xor_2str(and_2str(x, y), and_2str(x, z)), and_2str(y, z))
}

// this function takes a binary string and will right shift it by n bits
func shift_right(bit_string string, n int) string {
	res := bit_string[0 : len(bit_string)-n]
	return zeros_string(n) + res
}

// this function takes n as input ans returns a string of n zeros
func zeros_string(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "0"
	}
	return s
}

func s_0(x string) string {
	return xor_2str(xor_2str(rotate_right(x, 7), rotate_right(x, 18)), shift_right(x, 3))
}

func s_1(x string) string {
	return xor_2str(xor_2str(rotate_right(x, 17), rotate_right(x, 19)), shift_right(x, 10))
}

func message_schedule(i int, w_t []string) string {
	list := []int{
		int(dec_bin(s_1(w_t[i-2]))),
		int(dec_bin(w_t[i-7])),
		int(dec_bin(s_0(w_t[i-15]))),
		int(dec_bin(w_t[i-16])),
	}
	new_word := bin_32bit(int64(mod_32_addition(list)))
	return new_word
}

func mod_32_addition(input_set []int) int {
	value := 0
	for i := 0; i < len(input_set); i++ {
		value += input_set[i]
	}
	mod_32 := 4294967296
	return value % mod_32
}

func dec_bin(h string) int64 {
	n, _ := strconv.ParseInt(h, 2, 64)
	return n
}

func sha_256(input_string string) string {
	w_t := message_parsing(input_string)
	// fmt.Println(w_t)
	a := bin_32bit(dec_bin_hex(initial_hash_values[0]))
	b := bin_32bit(dec_bin_hex(initial_hash_values[1]))
	c := bin_32bit(dec_bin_hex(initial_hash_values[2]))
	d := bin_32bit(dec_bin_hex(initial_hash_values[3]))
	e := bin_32bit(dec_bin_hex(initial_hash_values[4]))
	f := bin_32bit(dec_bin_hex(initial_hash_values[5]))
	g := bin_32bit(dec_bin_hex(initial_hash_values[6]))
	h := bin_32bit(dec_bin_hex(initial_hash_values[7]))

	for i := 0; i < 64; i++ {

		if i > 15 {
			w_t = append(w_t, message_schedule(i, w_t))
		}
		list1 := []int{
			int(dec_bin(h)),
			int(dec_bin(e_1(e))),
			int(dec_bin(Ch(e, f, g))),
			int(dec_bin_hex(sha_256_constants[i])),
			int(dec_bin(w_t[i])),
		}
		t_1 := mod_32_addition(list1)
		list2 := []int{
			int(dec_bin(e_0(a))),
			int(dec_bin(Maj(a, b, c))),
		}
		t_2 := mod_32_addition(list2)
		h = g
		g = f
		f = e
		e = bin_32bit(int64(mod_32_addition([]int{int(dec_bin(d)), t_1})))
		d = c
		c = b
		b = a
		a = bin_32bit(int64(mod_32_addition([]int{t_1, t_2})))
	}

	hash_0 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[0])), int(dec_bin(a))})
	hash_1 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[1])), int(dec_bin(b))})
	hash_2 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[2])), int(dec_bin(c))})
	hash_3 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[3])), int(dec_bin(d))})
	hash_4 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[4])), int(dec_bin(e))})
	hash_5 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[5])), int(dec_bin(f))})
	hash_6 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[6])), int(dec_bin(g))})
	hash_7 := mod_32_addition([]int{int(dec_bin_hex(initial_hash_values[7])), int(dec_bin(h))})
	digest := hex_return(int64(hash_0)) + hex_return(int64(hash_1)) +
		hex_return(int64(hash_2)) + hex_return(int64(hash_3)) + hex_return(int64(hash_4)) +
		hex_return(int64(hash_5)) + hex_return(int64(hash_6)) + hex_return(int64(hash_7))
	return digest
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the input: ")
	s, _ := reader.ReadString('\n')
	s1 := ""
	for i := 0; i < len(s)-1; i++ {
		s1 = s1 + string(s[i])
	}
	fmt.Println("Hash value: ")
	fmt.Println(sha_256(s1))
}
