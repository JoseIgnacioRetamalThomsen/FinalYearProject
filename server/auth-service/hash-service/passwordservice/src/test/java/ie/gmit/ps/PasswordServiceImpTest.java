package ie.gmit.ps;

import static org.junit.Assert.*;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

import com.google.protobuf.BoolValue;
import com.google.protobuf.ByteString;

import io.grpc.stub.StreamObserver;

public class PasswordServiceImpTest {

    private static PasswordServiceImp ps = new PasswordServiceImp();

    private static final int PASSWORD_LEN = 32;
    private static final int SALT_LEN = 32;

    private final int userId = 123;
    private final String password = "password";

    private static int idResult = 0;
    private byte[] hashedPasswordResult;
    private byte[] saltResult;

    @Test
    public void testHash() {

        StreamObserver<HashResponse> responseObserver = new StreamObserver<HashResponse>() {

            @Override
            public void onNext(HashResponse value) {
                // TODO Auto-generated method stub
                idResult = value.getUserId();
                hashedPasswordResult = value.getHashedPassword().toByteArray();
                saltResult = value.getSalt().toByteArray();
            }

            @Override
            public void onError(Throwable t) {
            }

            @Override
            public void onCompleted() {
            }

        };

        HashRequest request = HashRequest.newBuilder().setUserId(userId).setPassword(password)
                    .build();

        ps.hash(request, responseObserver);

        assertEquals("Id", userId, idResult);
        assertEquals("Hashed pasworld lenght", PASSWORD_LEN, hashedPasswordResult.length);
        assertEquals("Salt lenght", SALT_LEN, saltResult.length);

    }

    private static final String shortPassword = "short";

    @Test(expected = io.grpc.StatusRuntimeException.class)
    public void testHashShortPassg() throws RuntimeException {
        StreamObserver<HashResponse> responseObserver = new StreamObserver<HashResponse>() {

            @Override
            public void onNext(HashResponse value) {
                // TODO Auto-generated method stub
                idResult = value.getUserId();
                hashedPasswordResult = value.getHashedPassword().toByteArray();
                saltResult = value.getSalt().toByteArray();
            }

            @Override
            public void onError(Throwable t) {
                throw (RuntimeException) t;
            }

            @Override
            public void onCompleted() {
            }

        };

        HashRequest request = HashRequest.newBuilder().setUserId(userId).setPassword(shortPassword)
                    .build();

        ps.hash(request, responseObserver);

    }

    private static final String longPassword = "llllllllllllllooooooooooooooooooonnnnnnnnnnnnnnnnggggggggggggggggggg";

    @Test(expected = io.grpc.StatusRuntimeException.class)
    public void testHashPasswordToLong() throws RuntimeException {
        StreamObserver<HashResponse> responseObserver = new StreamObserver<HashResponse>() {

            @Override
            public void onNext(HashResponse value) {
                // TODO Auto-generated method stub
                idResult = value.getUserId();
                hashedPasswordResult = value.getHashedPassword().toByteArray();
                saltResult = value.getSalt().toByteArray();
            }

            @Override
            public void onError(Throwable t) {
                throw (RuntimeException) t;
            }

            @Override
            public void onCompleted() {
            }

        };

        HashRequest request = HashRequest.newBuilder().setUserId(userId).setPassword(shortPassword)
                    .build();

        ps.hash(request, responseObserver);

    }

    private static int badId = -1;

    @Test(expected = io.grpc.StatusRuntimeException.class)
    public void testHashNegativeId() throws RuntimeException {
        StreamObserver<HashResponse> responseObserver = new StreamObserver<HashResponse>() {

            @Override
            public void onNext(HashResponse value) {
                // TODO Auto-generated method stub
                idResult = value.getUserId();
                hashedPasswordResult = value.getHashedPassword().toByteArray();
                saltResult = value.getSalt().toByteArray();
            }

            @Override
            public void onError(Throwable t) {
                throw (RuntimeException) t;
            }

            @Override
            public void onCompleted() {
            }

        };

        HashRequest request = HashRequest.newBuilder().setUserId(badId).setPassword(password)
                    .build();

        ps.hash(request, responseObserver);

    }

    private static boolean boolResponse;

    private final byte[] hashPassword = { 26, 21, 91, 62, -85, -19, 87, 8, 33, 35, -114, -63, -64,
            33, -111, -9, -50, 81, -115, -64, -40, 26, -46, -64, 115, -13, 40, -79, -119, 122, 115,
            -109 };
    private final byte[] salt = { 110, 53, -42, 49, 18, 25, -53, 121, -83, -83, 11, 64, 26, 118,
            -12, 51, -127, -108, -89, -6, 22, 23, 101, 45, 35, -56, -56, 53, -8, 126, -32, 81 };
    private final byte[] brongSalt = { 39, 123, -115, 95, 118, 41, 52, -99, 51, -14, -94, 11, 40,
            -99, -50, 26 };

    @Test
    public void testValidate() {
        StreamObserver<BoolValue> responseObserver = new StreamObserver<BoolValue>() {

            @Override
            public void onNext(BoolValue value) {

                boolResponse = value.getValue();
            }

            @Override
            public void onError(Throwable t) {
            }

            @Override
            public void onCompleted() {
            }

        };

        ValidateRequest request = ValidateRequest.newBuilder().setPassword(password)
                    .setHasshedPassword(ByteString.copyFrom(hashPassword)).setSalt(ByteString
                                .copyFrom(salt)).build();

        ps.validate(request, responseObserver);

        assertTrue(boolResponse);

    }
}
