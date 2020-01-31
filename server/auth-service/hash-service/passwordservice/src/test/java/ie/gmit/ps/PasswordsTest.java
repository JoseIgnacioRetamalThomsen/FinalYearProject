package ie.gmit.ps;

import static org.junit.Assert.*;

import java.util.Arrays;
import java.util.Comparator;
import java.util.HashSet;
import java.util.Set;

import org.junit.Test;

import ie.gmit.pshelper.Passwords;

public class PasswordsTest {

	private static final int expectedSaltLenght = 32;

	@Test
	public void testSaltLenght() {

		byte[] salt = Passwords.getNextSalt();

		int lenght = salt.length;

		assertEquals("Salt lenght is 32 bytes", expectedSaltLenght, lenght);
	}

	@Test
	public void testRandomSalt() {

		HashSet<Byte> set = new HashSet<>();

		for (int i = 0; i < 1000; i++) {
			assertTrue(set.add(new Byte(Passwords.getNextSalt())));
		}

	}

	@Test
	public void testHash() {
		byte[] salt = { 39, 123, -115, 95, 118, 41, 52, -99, 51, -14, -94, 11, 40, -99, -50, 26 };
		char[] password = "password".toCharArray();
		byte[] hashedPassword = Passwords.hash(password, salt);
		byte[] hashedPassword2 = Passwords.hash(password, salt);

		for (byte b : hashedPassword)
			System.out.print(b + ", ");
		assertNotNull(hashedPassword);

		assertArrayEquals("Generated paswords with same salt and passwords are equal", hashedPassword, hashedPassword2);

	}

	@Test
	public void testIsExpectedPassword() {
		byte[] salt = { 39, 123, -115, 95, 118, 41, 52, -99, 51, -14, -94, 11, 40, -99, -50, 26 };
		char[] password = "password".toCharArray();
		byte[] hashedPassword = { -101, -36, 95, -18, -57, -95, -110, -41, -62, 6, 26, 102, -62, 58, -1, -78, 80, 46,
				-17, -36, -35, -27, 42, 14, -4, -102, 68, -114, -113, -70, 69, 83 };

		assertTrue("Is Expected password", Passwords.isExpectedPassword(password, salt, hashedPassword));
	}
	
	@Test 
	public void testGenerateRandomPasswordLen() {
		int expectedLen = 16;
		String password = Passwords.generateRandomPassword(expectedLen);
		
		assertEquals("Generated password is of the requerid lenght", expectedLen,password.length());
		
	}
	
	@Test 
	public void testGenerateRandomPassword() {
		HashSet<String> set = new HashSet<>();
		
		for(int i =0 ;i<1000;i++) {
			assertTrue(set.add(Passwords.generateRandomPassword(10)));
		}
		
	}

}

class Byte {
	byte[] bytes;

	public Byte(byte[] b) {
		bytes = b;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + Arrays.hashCode(bytes);
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Byte other = (Byte) obj;
		if (!Arrays.equals(bytes, other.bytes))
			return false;
		return true;
	}

}