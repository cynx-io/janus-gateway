package com.plato

import com.core.GenericResponse
import com.plato.PlatoAnswerServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for plato.PlatoAnswerService.
 */
public object PlatoAnswerServiceGrpcKt {
  public const val SERVICE_NAME: String = PlatoAnswerServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val getAnswerByIdMethod: MethodDescriptor<AnswerIdRequest, AnswerResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getGetAnswerByIdMethod()

  public val getDetailAnswerByIdMethod: MethodDescriptor<AnswerIdRequest, DetailAnswerResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getGetDetailAnswerByIdMethod()

  public val listAnswersByTopicIdMethod: MethodDescriptor<TopicIdRequest, ListAnswersResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getListAnswersByTopicIdMethod()

  public val insertAnswerMethod: MethodDescriptor<InsertAnswerRequest, AnswerResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getInsertAnswerMethod()

  public val updateAnswerMethod: MethodDescriptor<UpdateAnswerRequest, AnswerResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getUpdateAnswerMethod()

  public val deleteAnswerMethod: MethodDescriptor<AnswerIdRequest, GenericResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getDeleteAnswerMethod()

  public val listDetailAnswersByTopicModeIdMethod:
      MethodDescriptor<TopicModeRequest, ListDetailAnswersResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getListDetailAnswersByTopicModeIdMethod()

  public val searchAnswersMethod: MethodDescriptor<SearchAnswersRequest, ListAnswersResponse>
    @JvmStatic
    get() = PlatoAnswerServiceGrpc.getSearchAnswersMethod()

  /**
   * A stub for issuing RPCs to a(n) plato.PlatoAnswerService service as suspending coroutines.
   */
  @StubFor(PlatoAnswerServiceGrpc::class)
  public class PlatoAnswerServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<PlatoAnswerServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): PlatoAnswerServiceCoroutineStub
        = PlatoAnswerServiceCoroutineStub(channel, callOptions)

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
    public suspend fun getAnswerById(request: AnswerIdRequest, headers: Metadata = Metadata()):
        AnswerResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getGetAnswerByIdMethod(),
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
    public suspend fun getDetailAnswerById(request: AnswerIdRequest, headers: Metadata =
        Metadata()): DetailAnswerResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getGetDetailAnswerByIdMethod(),
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
    public suspend fun listAnswersByTopicId(request: TopicIdRequest, headers: Metadata =
        Metadata()): ListAnswersResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getListAnswersByTopicIdMethod(),
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
    public suspend fun insertAnswer(request: InsertAnswerRequest, headers: Metadata = Metadata()):
        AnswerResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getInsertAnswerMethod(),
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
    public suspend fun updateAnswer(request: UpdateAnswerRequest, headers: Metadata = Metadata()):
        AnswerResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getUpdateAnswerMethod(),
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
    public suspend fun deleteAnswer(request: AnswerIdRequest, headers: Metadata = Metadata()):
        GenericResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getDeleteAnswerMethod(),
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
    public suspend fun listDetailAnswersByTopicModeId(request: TopicModeRequest, headers: Metadata =
        Metadata()): ListDetailAnswersResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getListDetailAnswersByTopicModeIdMethod(),
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
    public suspend fun searchAnswers(request: SearchAnswersRequest, headers: Metadata = Metadata()):
        ListAnswersResponse = unaryRpc(
      channel,
      PlatoAnswerServiceGrpc.getSearchAnswersMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the plato.PlatoAnswerService service based on Kotlin coroutines.
   */
  public abstract class PlatoAnswerServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.GetAnswerById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getAnswerById(request: AnswerIdRequest): AnswerResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.GetAnswerById is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.GetDetailAnswerById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getDetailAnswerById(request: AnswerIdRequest): DetailAnswerResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.GetDetailAnswerById is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.ListAnswersByTopicId.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listAnswersByTopicId(request: TopicIdRequest): ListAnswersResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.ListAnswersByTopicId is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.InsertAnswer.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun insertAnswer(request: InsertAnswerRequest): AnswerResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.InsertAnswer is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.UpdateAnswer.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateAnswer(request: UpdateAnswerRequest): AnswerResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.UpdateAnswer is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.DeleteAnswer.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteAnswer(request: AnswerIdRequest): GenericResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.DeleteAnswer is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.ListDetailAnswersByTopicModeId.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listDetailAnswersByTopicModeId(request: TopicModeRequest):
        ListDetailAnswersResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.ListDetailAnswersByTopicModeId is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerService.SearchAnswers.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun searchAnswers(request: SearchAnswersRequest): ListAnswersResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerService.SearchAnswers is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getGetAnswerByIdMethod(),
      implementation = ::getAnswerById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getGetDetailAnswerByIdMethod(),
      implementation = ::getDetailAnswerById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getListAnswersByTopicIdMethod(),
      implementation = ::listAnswersByTopicId
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getInsertAnswerMethod(),
      implementation = ::insertAnswer
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getUpdateAnswerMethod(),
      implementation = ::updateAnswer
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getDeleteAnswerMethod(),
      implementation = ::deleteAnswer
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getListDetailAnswersByTopicModeIdMethod(),
      implementation = ::listDetailAnswersByTopicModeId
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerServiceGrpc.getSearchAnswersMethod(),
      implementation = ::searchAnswers
    )).build()
  }
}
