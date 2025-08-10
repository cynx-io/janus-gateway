package com.plato

import com.core.GenericRequest
import com.core.GenericResponse
import com.plato.PlatoTopicServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for plato.PlatoTopicService.
 */
public object PlatoTopicServiceGrpcKt {
  public const val SERVICE_NAME: String = PlatoTopicServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val insertTopicMethod: MethodDescriptor<InsertTopicRequest, TopicResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getInsertTopicMethod()

  public val updateTopicMethod: MethodDescriptor<UpdateTopicRequest, TopicResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getUpdateTopicMethod()

  public val deleteTopicMethod: MethodDescriptor<TopicIdRequest, GenericResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getDeleteTopicMethod()

  public val listTopicsByUserIdMethod: MethodDescriptor<GenericRequest, PaginateTopicResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getListTopicsByUserIdMethod()

  public val paginateTopicMethod: MethodDescriptor<PaginateRequest, PaginateTopicResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getPaginateTopicMethod()

  public val getTopicByIdMethod: MethodDescriptor<TopicIdRequest, TopicResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getGetTopicByIdMethod()

  public val getTopicBySlugMethod: MethodDescriptor<SlugRequest, TopicResponse>
    @JvmStatic
    get() = PlatoTopicServiceGrpc.getGetTopicBySlugMethod()

  /**
   * A stub for issuing RPCs to a(n) plato.PlatoTopicService service as suspending coroutines.
   */
  @StubFor(PlatoTopicServiceGrpc::class)
  public class PlatoTopicServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<PlatoTopicServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): PlatoTopicServiceCoroutineStub =
        PlatoTopicServiceCoroutineStub(channel, callOptions)

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
    public suspend fun insertTopic(request: InsertTopicRequest, headers: Metadata = Metadata()):
        TopicResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getInsertTopicMethod(),
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
    public suspend fun updateTopic(request: UpdateTopicRequest, headers: Metadata = Metadata()):
        TopicResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getUpdateTopicMethod(),
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
    public suspend fun deleteTopic(request: TopicIdRequest, headers: Metadata = Metadata()):
        GenericResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getDeleteTopicMethod(),
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
    public suspend fun listTopicsByUserId(request: GenericRequest, headers: Metadata = Metadata()):
        PaginateTopicResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getListTopicsByUserIdMethod(),
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
    public suspend fun paginateTopic(request: PaginateRequest, headers: Metadata = Metadata()):
        PaginateTopicResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getPaginateTopicMethod(),
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
    public suspend fun getTopicById(request: TopicIdRequest, headers: Metadata = Metadata()):
        TopicResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getGetTopicByIdMethod(),
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
    public suspend fun getTopicBySlug(request: SlugRequest, headers: Metadata = Metadata()):
        TopicResponse = unaryRpc(
      channel,
      PlatoTopicServiceGrpc.getGetTopicBySlugMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the plato.PlatoTopicService service based on Kotlin coroutines.
   */
  public abstract class PlatoTopicServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for plato.PlatoTopicService.InsertTopic.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun insertTopic(request: InsertTopicRequest): TopicResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.InsertTopic is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoTopicService.UpdateTopic.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateTopic(request: UpdateTopicRequest): TopicResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.UpdateTopic is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoTopicService.DeleteTopic.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteTopic(request: TopicIdRequest): GenericResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.DeleteTopic is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoTopicService.ListTopicsByUserId.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listTopicsByUserId(request: GenericRequest): PaginateTopicResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.ListTopicsByUserId is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoTopicService.PaginateTopic.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun paginateTopic(request: PaginateRequest): PaginateTopicResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.PaginateTopic is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoTopicService.GetTopicById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getTopicById(request: TopicIdRequest): TopicResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.GetTopicById is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoTopicService.GetTopicBySlug.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getTopicBySlug(request: SlugRequest): TopicResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoTopicService.GetTopicBySlug is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getInsertTopicMethod(),
      implementation = ::insertTopic
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getUpdateTopicMethod(),
      implementation = ::updateTopic
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getDeleteTopicMethod(),
      implementation = ::deleteTopic
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getListTopicsByUserIdMethod(),
      implementation = ::listTopicsByUserId
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getPaginateTopicMethod(),
      implementation = ::paginateTopic
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getGetTopicByIdMethod(),
      implementation = ::getTopicById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoTopicServiceGrpc.getGetTopicBySlugMethod(),
      implementation = ::getTopicBySlug
    )).build()
  }
}
