/*
 * Jose Retamal
 * Distributed System Project
 * Galway-Mayo Institute of Technologies, 2019.
 */
package ie.gmit.ps;

import java.security.NoSuchAlgorithmException;
import java.security.spec.InvalidKeySpecException;
import java.util.logging.Logger;

import com.google.rpc.Status;
import ie.gmit.ps.PasswordServiceGrpc.PasswordServiceImplBase;
import ie.gmit.pshelper.Passwords;
import io.grpc.protobuf.StatusProto;
import io.grpc.stub.StreamObserver;
import io.grpc.Context;


import com.google.protobuf.BoolValue;
import com.google.protobuf.ByteString;
import com.google.rpc.Code;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.logging.Level;
import java.util.logging.Logger;


import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/***
 *
 * Server class, implements methods for hashing password using salt. Implements
 * two methods : {@code hash()} and {@code validate()}.
 *
 * @author Jose Retamal
 *
 */
public class PasswordServiceImp extends PasswordServiceImplBase {

    private final static int PASSWORD_MIN_LEN = 6;
    private final static int PASSWORD_MAX_LEN = 32;
    private final static int SALT_LEN = 32;
    private final static int HASHED_PAS_LEN = 32;

    private static final int SECONDS_TO_WAIT = 2;

    private static final ExecutorService CANCELLATION_EXECUTOR = Executors.newCachedThreadPool();

    private static final Logger logger = Logger.getLogger(PasswordServiceImp.class.getName());

    /***
     *
     */
    public void hash(HashRequest request, StreamObserver<HashResponse> responseObserver) {


        try {
/*
            Context context = Context.current();

            context.addListener(new Context.CancellationListener() {
                @Override
                public void cancelled(Context context) {
                    // CancellationCause is TimeoutException if it was exceeding the deadline
                    logger.log(Level.INFO, "deadlineExceeded(): The call was cancelled.", context.cancellationCause());

                }
            }, CANCELLATION_EXECUTOR);

            context.run(() -> {
                int secondsElapsed = 0;
                while (secondsElapsed < SECONDS_TO_WAIT && !context.isCancelled()) {
                    try {
                        Thread.sleep(1000L);
                    } catch (InterruptedException e) {
                    }
                    secondsElapsed++;
                }

                logger.log(Level.INFO, "deadlineExceeded(): The call ended after ~" + secondsElapsed + " seconds");
                throw new RuntimeException();

            });
*/
            byte[] salt = Passwords.getNextSalt();
            char[] password = request.getPassword().toCharArray();
            int id = request.getUserId();


            logger.info("New hash request, userId: " + request.getUserId());

            if (id < 0 || password.length < PASSWORD_MIN_LEN
                    || password.length > PASSWORD_MAX_LEN) {

                throw new InvalidRequestException();

            }

            byte[] hashed = Passwords.hash(password, salt);

            responseObserver.onNext(HashResponse.newBuilder().setUserId(id).setHashedPassword(
                    ByteString.copyFrom(hashed)).setSalt(ByteString.copyFrom(salt)).build());

            responseObserver.onCompleted();

        } catch (InvalidRequestException ex) {

            Status status = Status.newBuilder().setCode(Code.INVALID_ARGUMENT.getNumber())
                    .setMessage("Id or password malformed").build();
            logger.info("Inalid request ");

            responseObserver.onError(StatusProto.toStatusRuntimeException(status));

        } catch (AssertionError err) {
            Status status = Status.newBuilder().setCode(Code.INTERNAL.getNumber())
                    .setMessage("Fail to hash").build();
            logger.info("FAil to hash ");

            responseObserver.onError(StatusProto.toStatusRuntimeException(status));
        } catch (RuntimeException ex) {

            responseObserver.onError(ex);

        }

    }

    @Override
    public void validate(ValidateRequest request,
                         io.grpc.stub.StreamObserver<BoolValue> responseObserver) {

        boolean isValid = false;

        logger.info("New validate request.");

        try {

            char[] password = request.getPassword().toCharArray();
            byte[] salt = request.getSalt().toByteArray();
            byte[] hashedPassword = request.getHasshedPassword().toByteArray();

            if (password.length < PASSWORD_MIN_LEN || password.length > PASSWORD_MAX_LEN
                    || salt.length != SALT_LEN || hashedPassword.length != HASHED_PAS_LEN) {

                throw new InvalidRequestException();

            }

            isValid = SlowEquals.isExpectedPassword(password, salt, hashedPassword);

            responseObserver.onNext(BoolValue.newBuilder().setValue(isValid).build());

        } catch (InvalidRequestException ex) {

            Status status = Status.newBuilder().setCode(Code.INVALID_ARGUMENT.getNumber())
                    .setMessage("Password or salt or hashedPassword malformed ").build();

            logger.info("Inalid request ");

            responseObserver.onError(StatusProto.toStatusRuntimeException(status));

        } catch (RuntimeException ex) {

            responseObserver.onError(ex);
        }

        responseObserver.onCompleted();
    }

}
