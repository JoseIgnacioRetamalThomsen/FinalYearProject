package ie.gmit.ps;

import ie.gmit.pshelper.Passwords;

/**
 * Implements slow equals function, which compare two arrays of bytes in
 * constant time, using {@code Passwords} class.
 * 
 * @author Jose I. Retamal
 *
 */
public class SlowEquals {

	/**
	 * Returns true if the given password and salt match the hashed value, false
	 * otherwise.<br>
	 *
	 * @param password     the password to check
	 * @param salt         the salt used to hash the password
	 * @param expectedHash the expected hashed value of the password
	 * @return true if the given password and salt match the hashed value, false
	 *         otherwise
	 */
	public static boolean isExpectedPassword(char[] password, byte[] salt, byte[] expectedHash) {
		byte[] pwdHash = Passwords.hash(password, salt);
		return slowEquals(pwdHash, expectedHash);
	}

	/**
	 * 
	 * Compare 2 byte arrays in constant time, Adapted from
	 * <a href="https://crackstation.net/hashing-security.htm">
	 * 
	 * @param a array of bytes
	 * @param b array of bytes
	 * @return true if the arrays are equals
	 */
	private static boolean slowEquals(byte[] a, byte[] b) {
		int diff = a.length ^ b.length;
		for (int i = 0; i < a.length && i < b.length; i++)
			diff |= a[i] ^ b[i];
		return diff == 0;
	}
}
