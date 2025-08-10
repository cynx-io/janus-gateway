package com.plato

import com.core.GenericResponse
import com.plato.PlatoModeServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for plato.PlatoModeService.
 */
public object PlatoModeServiceGrpcKt {
  public const val SERVICE_NAME: String = PlatoModeServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val getModeByIdMethod: MethodDescriptor<ModeIdRequest, ModeResponse>
    @JvmStatic
    get() = PlatoModeServiceGrpc.getGetModeByIdMethod()

  public val insertModeMethod: MethodDescriptor<InsertModeRequest, ModeResponse>
    @JvmStatic
    get() = PlatoModeServiceGrpc.getInsertModeMethod()

  public val updateModeMethod: MethodDescriptor<UpdateModeRequest, ModeResponse>
    @JvmStatic
    get() = PlatoModeServiceGrpc.getUpdateModeMethod()

  public val deleteModeMethod: MethodDescriptor<ModeIdRequest, GenericResponse>
    @JvmStatic
    get() = PlatoModeServiceGrpc.getDeleteModeMethod()

  public val listModesByTopicIdMethod: MethodDescriptor<TopicIdRequest, ListModesResponse>
    @JvmStatic
    get() = PlatoModeServiceGrpc.getListModesByTopicIdMethod()

  /**
   * A stub for issuing RPCs to a(n) plato.PlatoModeService service as suspending coroutines.
   */
  @StubFor(PlatoModeServiceGrpc::class)
  public class PlatoModeServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<PlatoModeServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): PlatoModeServiceCoroutineStub =
        PlatoModeServiceCoroutineStub(channel, callOptions)

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
    public suspend fun getModeById(request: ModeIdRequest, headers: Metadata = Metadata()):
        ModeResponse = unaryRpc(
      channel,
      PlatoModeServiceGrpc.getGetModeByIdMethod(),
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
    public suspend fun insertMode(request: InsertModeRequest, headers: Metadata = Metadata()):
        ModeResponse = unaryRpc(
      channel,
      PlatoModeServiceGrpc.getInsertModeMethod(),
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
    public suspend fun updateMode(request: UpdateModeRequest, headers: Metadata = Metadata()):
        ModeResponse = unaryRpc(
      channel,
      PlatoModeServiceGrpc.getUpdateModeMethod(),
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
    public suspend fun deleteMode(request: ModeIdRequest, headers: Metadata = Metadata()):
        GenericResponse = unaryRpc(
      channel,
      PlatoModeServiceGrpc.getDeleteModeMethod(),
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
    public suspend fun listModesByTopicId(request: TopicIdRequest, headers: Metadata = Metadata()):
        ListModesResponse = unaryRpc(
      channel,
      PlatoModeServiceGrpc.getListModesByTopicIdMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the plato.PlatoModeService service based on Kotlin coroutines.
   */
  public abstract class PlatoModeServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for plato.PlatoModeService.GetModeById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getModeById(request: ModeIdRequest): ModeResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoModeService.GetModeById is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoModeService.InsertMode.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun insertMode(request: InsertModeRequest): ModeResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoModeService.InsertMode is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoModeService.UpdateMode.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateMode(request: UpdateModeRequest): ModeResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoModeService.UpdateMode is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoModeService.DeleteMode.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteMode(request: ModeIdRequest): GenericResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoModeService.DeleteMode is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoModeService.ListModesByTopicId.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listModesByTopicId(request: TopicIdRequest): ListModesResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoModeService.ListModesByTopicId is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoModeServiceGrpc.getGetModeByIdMethod(),
      implementation = ::getModeById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoModeServiceGrpc.getInsertModeMethod(),
      implementation = ::insertMode
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoModeServiceGrpc.getUpdateModeMethod(),
      implementation = ::updateMode
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoModeServiceGrpc.getDeleteModeMethod(),
      implementation = ::deleteMode
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoModeServiceGrpc.getListModesByTopicIdMethod(),
      implementation = ::listModesByTopicId
    )).build()
  }
}
