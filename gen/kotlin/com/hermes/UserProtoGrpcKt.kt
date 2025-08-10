package com.hermes

import com.core.GenericRequest
import com.hermes.HermesUserServiceGrpc.getServiceDescriptor
import io.grpc.CallOptions
import io.grpc.CallOptions.DEFAULT
import io.grpc.Channel
import io.grpc.Metadata
import io.grpc.MethodDescriptor
import io.grpc.ServerServiceDefinition
import io.grpc.ServerServiceDefinition.builder
import io.grpc.ServiceDescriptor
import io.grpc.Status.UNIMPLEMENTED
import io.grpc.StatusException
import io.grpc.kotlin.AbstractCoroutineServerImpl
import io.grpc.kotlin.AbstractCoroutineStub
import io.grpc.kotlin.ClientCalls.unaryRpc
import io.grpc.kotlin.ServerCalls.unaryServerMethodDefinition
import io.grpc.kotlin.StubFor
import kotlin.String
import kotlin.coroutines.CoroutineContext
import kotlin.coroutines.EmptyCoroutineContext
import kotlin.jvm.JvmOverloads
import kotlin.jvm.JvmStatic

/**
 * Holder for Kotlin coroutine-based client and server APIs for hermes.HermesUserService.
 */
public object HermesUserServiceGrpcKt {
  public const val SERVICE_NAME: String = HermesUserServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val checkUsernameMethod: MethodDescriptor<UsernameRequest, CheckUsernameResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getCheckUsernameMethod()

  public val getUserMethod: MethodDescriptor<UsernameRequest, UserResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getGetUserMethod()

  public val createUserMethod: MethodDescriptor<UsernamePasswordRequest, UserResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getCreateUserMethod()

  public val createUserFromGuestMethod: MethodDescriptor<UsernamePasswordRequest, UserResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getCreateUserFromGuestMethod()

  public val upsertGuestUserMethod: MethodDescriptor<GenericRequest, UserResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getUpsertGuestUserMethod()

  public val paginateUsersMethod: MethodDescriptor<PaginateRequest, PaginateUsersResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getPaginateUsersMethod()

  public val validatePasswordMethod: MethodDescriptor<UsernamePasswordRequest, UserResponse>
    @JvmStatic
    get() = HermesUserServiceGrpc.getValidatePasswordMethod()

  /**
   * A stub for issuing RPCs to a(n) hermes.HermesUserService service as suspending coroutines.
   */
  @StubFor(HermesUserServiceGrpc::class)
  public class HermesUserServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<HermesUserServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): HermesUserServiceCoroutineStub =
        HermesUserServiceCoroutineStub(channel, callOptions)

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun checkUsername(request: UsernameRequest, headers: Metadata = Metadata()):
        CheckUsernameResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getCheckUsernameMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun getUser(request: UsernameRequest, headers: Metadata = Metadata()):
        UserResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getGetUserMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun createUser(request: UsernamePasswordRequest, headers: Metadata = Metadata()):
        UserResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getCreateUserMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun createUserFromGuest(request: UsernamePasswordRequest, headers: Metadata =
        Metadata()): UserResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getCreateUserFromGuestMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun upsertGuestUser(request: GenericRequest, headers: Metadata = Metadata()):
        UserResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getUpsertGuestUserMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun paginateUsers(request: PaginateRequest, headers: Metadata = Metadata()):
        PaginateUsersResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getPaginateUsersMethod(),
      request,
      callOptions,
      headers
    )

    /**
     * Executes this RPC and returns the response message, suspending until the RPC completes
     * with [`Status.OK`][io.grpc.Status].  If the RPC completes with another status, a
     * corresponding
     * [StatusException] is thrown.  If this coroutine is cancelled, the RPC is also cancelled
     * with the corresponding exception as a cause.
     *
     * @param request The request message to send to the server.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return The single response from the server.
     */
    public suspend fun validatePassword(request: UsernamePasswordRequest, headers: Metadata =
        Metadata()): UserResponse = unaryRpc(
      channel,
      HermesUserServiceGrpc.getValidatePasswordMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the hermes.HermesUserService service based on Kotlin coroutines.
   */
  public abstract class HermesUserServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for hermes.HermesUserService.CheckUsername.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun checkUsername(request: UsernameRequest): CheckUsernameResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.CheckUsername is unimplemented"))

    /**
     * Returns the response to an RPC for hermes.HermesUserService.GetUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getUser(request: UsernameRequest): UserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.GetUser is unimplemented"))

    /**
     * Returns the response to an RPC for hermes.HermesUserService.CreateUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun createUser(request: UsernamePasswordRequest): UserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.CreateUser is unimplemented"))

    /**
     * Returns the response to an RPC for hermes.HermesUserService.CreateUserFromGuest.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun createUserFromGuest(request: UsernamePasswordRequest): UserResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.CreateUserFromGuest is unimplemented"))

    /**
     * Returns the response to an RPC for hermes.HermesUserService.UpsertGuestUser.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun upsertGuestUser(request: GenericRequest): UserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.UpsertGuestUser is unimplemented"))

    /**
     * Returns the response to an RPC for hermes.HermesUserService.PaginateUsers.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun paginateUsers(request: PaginateRequest): PaginateUsersResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.PaginateUsers is unimplemented"))

    /**
     * Returns the response to an RPC for hermes.HermesUserService.ValidatePassword.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun validatePassword(request: UsernamePasswordRequest): UserResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method hermes.HermesUserService.ValidatePassword is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getCheckUsernameMethod(),
      implementation = ::checkUsername
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getGetUserMethod(),
      implementation = ::getUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getCreateUserMethod(),
      implementation = ::createUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getCreateUserFromGuestMethod(),
      implementation = ::createUserFromGuest
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getUpsertGuestUserMethod(),
      implementation = ::upsertGuestUser
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getPaginateUsersMethod(),
      implementation = ::paginateUsers
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = HermesUserServiceGrpc.getValidatePasswordMethod(),
      implementation = ::validatePassword
    )).build()
  }
}
