package ie.gmit.ps;

import static org.junit.Assert.*;

import org.junit.Test;

public class SlowEqualsTest {

	private static final byte[] array1 = {-95, -110, -41, -62, 6, 26, 102, -62, 58, -1, -78, 80, 46} ;
	private static final byte[] array2 = {-95, -110, -41, -62, 6, 26, 102, -62, 58, -1, -78, 80, 46} ;
	private static final byte[] array3 = {-95, -110, -4, -62, 6, 26, 102, -62, 58, -1, -78, 80, 46} ;
	private static final byte[] array4 = {-9, -1} ;
	
	byte[] salt = { 39, 123, -115, 95, 118, 41, 52, -99, 51, -14, -94, 11, 40, -99, -50, 26 };
	char[] password = "password".toCharArray();
	byte[] hashedPassword = { -101, -36, 95, -18, -57, -95, -110, -41, -62, 6, 26, 102, -62, 58, -1, -78, 80, 46,
			-17, -36, -35, -27, 42, 14, -4, -102, 68, -114, -113, -70, 69, 83 };
	byte[] otherPassword = { -101, -36, 95, -18, -57, -95, -110, -41, -62, 6, 26, 102, -62, 58, -1, -78, 80, 46,
			-17, -36, -35, -27, 42, 14, -4, -102, 68, -114, -11, -70, 69, 83 };
	
	@Test
	public void testIsExpectedPassword() {
		assertTrue("",SlowEquals.isExpectedPassword(password, salt, hashedPassword));
		
		assertFalse("",SlowEquals.isExpectedPassword(password, salt, otherPassword));
	}

}
