package ie.gmit.wcity.profile;

import java.io.File;
import java.io.IOException;
import java.util.logging.Logger;

import io.grpc.Server;
import io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.NettyServerBuilder;
import io.netty.handler.ssl.ClientAuth;
import io.netty.handler.ssl.SslContextBuilder;


public class ProfileDBServer {

	 private Server server;

	    private static final Logger logger = Logger.getLogger(ProfileDBServer.class.getName());
	    private static final int DEFAULT_PORT = 5777;

	    private final int port;
	    private final String certChainFilePath;
	    private final String privateKeyFilePath;
	    private final String trustCertCollectionFilePath;

	    /***
	     *
	     *  Build the services.
	     *
	     * @param port Port number for run the service.
	     * @param certChainFilePath if is null will run without tsl
	     * @param privateKeyFilePath if is null will run without tsl
	     * @param trustCertCollectionFilePath provide this for mutual tsl ,if null will run just normal
	     */
	    public ProfileDBServer(int port,
	                          String certChainFilePath,
	                          String privateKeyFilePath,
	                          String trustCertCollectionFilePath) {

	        this.port = port;
	        this.certChainFilePath = certChainFilePath;
	        this.privateKeyFilePath = privateKeyFilePath;
	        this.trustCertCollectionFilePath = trustCertCollectionFilePath;
	    }

	    /***
	     * Build context using certificates.
	     *
	     * @return
	     */
	    private SslContextBuilder getSslContextBuilder() {

	        SslContextBuilder sslClientContextBuilder;



	            sslClientContextBuilder = SslContextBuilder.forServer(
	                    new File(certChainFilePath),
	                    new File(privateKeyFilePath));

	            if (trustCertCollectionFilePath != null) {
	                sslClientContextBuilder.trustManager(new File(trustCertCollectionFilePath));
	                sslClientContextBuilder.clientAuth(ClientAuth.REQUIRE);
	            }

	        return GrpcSslContexts.configure(sslClientContextBuilder);
	    }

	    /***
	     * Start the serve, will use context building if certificates are provided,
	     * if not will run with out authentication.
	     * Shutdown if VM is shutdown.
	     *
	     * @throws IOException
	     */
	    private void start() throws IOException {

	        if (privateKeyFilePath != null) {
	            SslContextBuilder sslContextBuilder = getSslContextBuilder();

	            server = NettyServerBuilder.forPort(port)
	                    .addService(new ConcurrentProfileDBImp())
	                    .sslContext(sslContextBuilder.build())
	                    .build()
	                    .start();

	        } else {

	            server = NettyServerBuilder.forPort(port)
	                    .addService(new ConcurrentProfileDBImp())
	                    .build()
	                    .start();

	        }

	        logger.info("Server started, listening on " + port);

	        Runtime.getRuntime().addShutdownHook(new Thread() {
	            @Override
	            public void run() {

	                // Use stderr here since the logger may have been reset by its JVM shutdown hook.
	                System.err.println("*** shutting down gRPC server since JVM is shutting down");
	                ProfileDBServer.this.stop();
	                System.err.println("*** server shut down");

	            }
	        });
	    }


	    /***
	     * Shutdown services.
	     */
	    private void stop() {

	        if (server != null) {

	            server.shutdown();

	        }
	    }

	    /**
	     * Await termination on the main thread since the grpc library uses daemon
	     * threads.
	     */
	    private void blockUntilShutdown() throws InterruptedException {

	        if (server != null) {

	            server.awaitTermination();

	        }
	    }

	    /***
	     * Create and start the services, if not args provided will run on default port,
	     * If not certificates provides will run with out tsl.
	     *
	     *
	     * @param args can have 0 args: run on default  port, 3 args: port certChainFilePath privateKeyFilePath,
	     *             or 4 arg: last is trustCertCollectionFilePath.
	     */
	    public static void main(String[] args) {

	        try {
	            if (args.length == 2 || args.length > 4) {
	                System.out.println(
	                        "USAGE WITH OUT TSL: " +
	                                "PasswordServer port \n if you don't supply a port it will run in port " + DEFAULT_PORT +
	                                ". \nUSAGE WITH TSL: " +
	                                "HelloWorldServerTls port certChainFilePath privateKeyFilePath " +
	                                "[trustCertCollectionFilePath]\n  Note: You only need to supply trustCertCollectionFilePath if you want " +
	                                "to enable Mutual TLS.");
	                System.exit(0);
	            }

	            final ProfileDBServer server = new ProfileDBServer(
	                    args.length == 0 ? DEFAULT_PORT : Integer.parseInt(args[0]),
	                    (args.length == 3 || args.length ==4) ? args[1] : null,
	                    (args.length == 3|| args.length ==4)  ? args[2] : null,
	                    args.length == 4 ? args[3] : null
	            );

	            server.start();

	            server.blockUntilShutdown();
	        } catch (IOException | InterruptedException ie) {
	            System.out.println("Server exception :" + ie);
	            System.exit(-1);
	        }
	    }
}
