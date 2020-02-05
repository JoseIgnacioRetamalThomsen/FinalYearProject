package ie.gmit.wcity.profile;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.logging.Logger;

import io.grpc.Context;
import io.grpc.stub.StreamObserver;
import io.grpc.wcity.profilesDB.CreateUserRequestPDB;
import io.grpc.wcity.profilesDB.CreateUserResponsePDB;
import io.grpc.wcity.profilesDB.ProfilesDBGrpc.ProfilesDBImplBase;

public class ProfileDBImp extends ProfilesDBImplBase {

	private final static String URL = "bolt://10.154.0.6:7687";
	private final static String USER_NAME = "neo4j";
	private final static String PASSWORD = "test";

	private static final ExecutorService CANCELLATION_EXECUTOR = Executors.newCachedThreadPool();

	private static final Logger logger = Logger.getLogger(ProfileDBImp.class.getName());
	
	public void createUser(CreateUserRequestPDB request, StreamObserver<CreateUserResponsePDB> response) {
		logger.info("Create user request.");
		try {
			Context context = Context.current();
			DAO dao = new DAO(URL, USER_NAME, PASSWORD);
			
			dao.AddUser(request.getEmail(), request.getName(), request.getDescription());
			
			response.onNext( CreateUserResponsePDB.newBuilder().setEmail(request.getEmail()).
					setValied("true")
					.build());
			 response.onCompleted();
		}catch(Exception e) {
			response.onError(e);
			
		}
	}

	/*    *//***
			*
			*//*
				 * public void hash(HashRequest request, StreamObserver<HashResponse>
				 * responseObserver) {
				 * 
				 * 
				 * try { /* Context context = Context.current();
				 * 
				 * context.addListener(new Context.CancellationListener() {
				 * 
				 * @Override public void cancelled(Context context) { // CancellationCause is
				 * TimeoutException if it was exceeding the deadline logger.log(Level.INFO,
				 * "deadlineExceeded(): The call was cancelled.", context.cancellationCause());
				 * 
				 * } }, CANCELLATION_EXECUTOR);
				 * 
				 * context.run(() -> { int secondsElapsed = 0; while (secondsElapsed <
				 * SECONDS_TO_WAIT && !context.isCancelled()) { try { Thread.sleep(1000L); }
				 * catch (InterruptedException e) { } secondsElapsed++; }
				 * 
				 * logger.log(Level.INFO, "deadlineExceeded(): The call ended after ~" +
				 * secondsElapsed + " seconds"); throw new RuntimeException();
				 * 
				 * });
				 * 
				 * byte[] salt = Passwords.getNextSalt(); char[] password =
				 * request.getPassword().toCharArray(); int id = request.getUserId();
				 * 
				 * 
				 * logger.info("New hash request, userId: " + request.getUserId());
				 * 
				 * if (id < 0 || password.length < PASSWORD_MIN_LEN || password.length >
				 * PASSWORD_MAX_LEN) {
				 * 
				 * throw new InvalidRequestException();
				 * 
				 * }
				 * 
				 * byte[] hashed = Passwords.hash(password, salt);
				 * 
				 * responseObserver.onNext(HashResponse.newBuilder().setUserId(id).
				 * setHashedPassword(
				 * ByteString.copyFrom(hashed)).setSalt(ByteString.copyFrom(salt)).build());
				 * 
				 * responseObserver.onCompleted();
				 * 
				 * } catch (InvalidRequestException ex) {
				 * 
				 * Status status =
				 * Status.newBuilder().setCode(Code.INVALID_ARGUMENT.getNumber())
				 * .setMessage("Id or password malformed").build();
				 * logger.info("Inalid request ");
				 * 
				 * responseObserver.onError(StatusProto.toStatusRuntimeException(status));
				 * 
				 * } catch (AssertionError err) { Status status =
				 * Status.newBuilder().setCode(Code.INTERNAL.getNumber())
				 * .setMessage("Fail to hash").build(); logger.info("FAil to hash ");
				 * 
				 * responseObserver.onError(StatusProto.toStatusRuntimeException(status)); }
				 * catch (RuntimeException ex) {
				 * 
				 * responseObserver.onError(ex);
				 * 
				 * }
				 * 
				 * }
				 * 
				 * @Override public void validate(ValidateRequest request,
				 * io.grpc.stub.StreamObserver<BoolValue> responseObserver) {
				 * 
				 * boolean isValid = false;
				 * 
				 * logger.info("New validate request.");
				 * 
				 * try {
				 * 
				 * char[] password = request.getPassword().toCharArray(); byte[] salt =
				 * request.getSalt().toByteArray(); byte[] hashedPassword =
				 * request.getHasshedPassword().toByteArray();
				 * 
				 * if (password.length < PASSWORD_MIN_LEN || password.length > PASSWORD_MAX_LEN
				 * || salt.length != SALT_LEN || hashedPassword.length != HASHED_PAS_LEN) {
				 * 
				 * throw new InvalidRequestException();
				 * 
				 * }
				 * 
				 * isValid = SlowEquals.isExpectedPassword(password, salt, hashedPassword);
				 * 
				 * responseObserver.onNext(BoolValue.newBuilder().setValue(isValid).build());
				 * 
				 * } catch (InvalidRequestException ex) {
				 * 
				 * Status status =
				 * Status.newBuilder().setCode(Code.INVALID_ARGUMENT.getNumber())
				 * .setMessage("Password or salt or hashedPassword malformed ").build();
				 * 
				 * logger.info("Inalid request ");
				 * 
				 * responseObserver.onError(StatusProto.toStatusRuntimeException(status));
				 * 
				 * } catch (RuntimeException ex) {
				 * 
				 * responseObserver.onError(ex); }
				 * 
				 * responseObserver.onCompleted(); }
				 */

}
