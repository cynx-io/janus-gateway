package com.plato

import com.plato.PlatoDailyGameServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for plato.PlatoDailyGameService.
 */
public object PlatoDailyGameServiceGrpcKt {
  public const val SERVICE_NAME: String = PlatoDailyGameServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val getDetailDailyGameByIdMethod:
      MethodDescriptor<DailyGameIdRequest, DetailDailyGameResponse>
    @JvmStatic
    get() = PlatoDailyGameServiceGrpc.getGetDetailDailyGameByIdMethod()

  public val getModeDailyGameByIdMethod: MethodDescriptor<ModeIdRequest, ModeDailyGameResponse>
    @JvmStatic
    get() = PlatoDailyGameServiceGrpc.getGetModeDailyGameByIdMethod()

  public val getPublicDailyGameMethod: MethodDescriptor<ModeIdRequest, PublicDailyGameResponse>
    @JvmStatic
    get() = PlatoDailyGameServiceGrpc.getGetPublicDailyGameMethod()

  public val attemptAnswerMethod: MethodDescriptor<AttemptAnswerRequest, AttemptAnswerResponse>
    @JvmStatic
    get() = PlatoDailyGameServiceGrpc.getAttemptAnswerMethod()

  public val attemptHistoryMethod: MethodDescriptor<DailyGameIdRequest, AttemptHistoryResponse>
    @JvmStatic
    get() = PlatoDailyGameServiceGrpc.getAttemptHistoryMethod()

  /**
   * A stub for issuing RPCs to a(n) plato.PlatoDailyGameService service as suspending coroutines.
   */
  @StubFor(PlatoDailyGameServiceGrpc::class)
  public class PlatoDailyGameServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<PlatoDailyGameServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions):
        PlatoDailyGameServiceCoroutineStub = PlatoDailyGameServiceCoroutineStub(channel,
        callOptions)

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
    public suspend fun getDetailDailyGameById(request: DailyGameIdRequest, headers: Metadata =
        Metadata()): DetailDailyGameResponse = unaryRpc(
      channel,
      PlatoDailyGameServiceGrpc.getGetDetailDailyGameByIdMethod(),
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
    public suspend fun getModeDailyGameById(request: ModeIdRequest, headers: Metadata = Metadata()):
        ModeDailyGameResponse = unaryRpc(
      channel,
      PlatoDailyGameServiceGrpc.getGetModeDailyGameByIdMethod(),
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
    public suspend fun getPublicDailyGame(request: ModeIdRequest, headers: Metadata = Metadata()):
        PublicDailyGameResponse = unaryRpc(
      channel,
      PlatoDailyGameServiceGrpc.getGetPublicDailyGameMethod(),
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
    public suspend fun attemptAnswer(request: AttemptAnswerRequest, headers: Metadata = Metadata()):
        AttemptAnswerResponse = unaryRpc(
      channel,
      PlatoDailyGameServiceGrpc.getAttemptAnswerMethod(),
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
    public suspend fun attemptHistory(request: DailyGameIdRequest, headers: Metadata = Metadata()):
        AttemptHistoryResponse = unaryRpc(
      channel,
      PlatoDailyGameServiceGrpc.getAttemptHistoryMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the plato.PlatoDailyGameService service based on Kotlin coroutines.
   */
  public abstract class PlatoDailyGameServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for plato.PlatoDailyGameService.GetDetailDailyGameById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getDetailDailyGameById(request: DailyGameIdRequest):
        DetailDailyGameResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoDailyGameService.GetDetailDailyGameById is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoDailyGameService.GetModeDailyGameById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getModeDailyGameById(request: ModeIdRequest): ModeDailyGameResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoDailyGameService.GetModeDailyGameById is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoDailyGameService.GetPublicDailyGame.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getPublicDailyGame(request: ModeIdRequest): PublicDailyGameResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoDailyGameService.GetPublicDailyGame is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoDailyGameService.AttemptAnswer.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun attemptAnswer(request: AttemptAnswerRequest): AttemptAnswerResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoDailyGameService.AttemptAnswer is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoDailyGameService.AttemptHistory.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun attemptHistory(request: DailyGameIdRequest): AttemptHistoryResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoDailyGameService.AttemptHistory is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoDailyGameServiceGrpc.getGetDetailDailyGameByIdMethod(),
      implementation = ::getDetailDailyGameById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoDailyGameServiceGrpc.getGetModeDailyGameByIdMethod(),
      implementation = ::getModeDailyGameById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoDailyGameServiceGrpc.getGetPublicDailyGameMethod(),
      implementation = ::getPublicDailyGame
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoDailyGameServiceGrpc.getAttemptAnswerMethod(),
      implementation = ::attemptAnswer
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoDailyGameServiceGrpc.getAttemptHistoryMethod(),
      implementation = ::attemptHistory
    )).build()
  }
}
