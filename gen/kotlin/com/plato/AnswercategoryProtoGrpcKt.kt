package com.plato

import com.core.GenericResponse
import com.plato.PlatoAnswerCategoryServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for plato.PlatoAnswerCategoryService.
 */
public object PlatoAnswerCategoryServiceGrpcKt {
  public const val SERVICE_NAME: String = PlatoAnswerCategoryServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val getAnswerCategoryByIdMethod:
      MethodDescriptor<AnswerCategoryIdRequest, AnswerCategoryResponse>
    @JvmStatic
    get() = PlatoAnswerCategoryServiceGrpc.getGetAnswerCategoryByIdMethod()

  public val listAnswerCategoriesByAnswerIdMethod:
      MethodDescriptor<AnswerIdRequest, ListAnswerCategoriesResponse>
    @JvmStatic
    get() = PlatoAnswerCategoryServiceGrpc.getListAnswerCategoriesByAnswerIdMethod()

  public val insertAnswerCategoryMethod:
      MethodDescriptor<InsertAnswerCategoryRequest, AnswerCategoryResponse>
    @JvmStatic
    get() = PlatoAnswerCategoryServiceGrpc.getInsertAnswerCategoryMethod()

  public val updateAnswerCategoryMethod:
      MethodDescriptor<UpdateAnswerCategoryRequest, AnswerCategoryResponse>
    @JvmStatic
    get() = PlatoAnswerCategoryServiceGrpc.getUpdateAnswerCategoryMethod()

  public val deleteAnswerCategoryMethod: MethodDescriptor<AnswerCategoryIdRequest, GenericResponse>
    @JvmStatic
    get() = PlatoAnswerCategoryServiceGrpc.getDeleteAnswerCategoryMethod()

  /**
   * A stub for issuing RPCs to a(n) plato.PlatoAnswerCategoryService service as suspending
   * coroutines.
   */
  @StubFor(PlatoAnswerCategoryServiceGrpc::class)
  public class PlatoAnswerCategoryServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<PlatoAnswerCategoryServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions):
        PlatoAnswerCategoryServiceCoroutineStub = PlatoAnswerCategoryServiceCoroutineStub(channel,
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
    public suspend fun getAnswerCategoryById(request: AnswerCategoryIdRequest, headers: Metadata =
        Metadata()): AnswerCategoryResponse = unaryRpc(
      channel,
      PlatoAnswerCategoryServiceGrpc.getGetAnswerCategoryByIdMethod(),
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
    public suspend fun listAnswerCategoriesByAnswerId(request: AnswerIdRequest, headers: Metadata =
        Metadata()): ListAnswerCategoriesResponse = unaryRpc(
      channel,
      PlatoAnswerCategoryServiceGrpc.getListAnswerCategoriesByAnswerIdMethod(),
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
    public suspend fun insertAnswerCategory(request: InsertAnswerCategoryRequest, headers: Metadata
        = Metadata()): AnswerCategoryResponse = unaryRpc(
      channel,
      PlatoAnswerCategoryServiceGrpc.getInsertAnswerCategoryMethod(),
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
    public suspend fun updateAnswerCategory(request: UpdateAnswerCategoryRequest, headers: Metadata
        = Metadata()): AnswerCategoryResponse = unaryRpc(
      channel,
      PlatoAnswerCategoryServiceGrpc.getUpdateAnswerCategoryMethod(),
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
    public suspend fun deleteAnswerCategory(request: AnswerCategoryIdRequest, headers: Metadata =
        Metadata()): GenericResponse = unaryRpc(
      channel,
      PlatoAnswerCategoryServiceGrpc.getDeleteAnswerCategoryMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the plato.PlatoAnswerCategoryService service based on Kotlin
   * coroutines.
   */
  public abstract class PlatoAnswerCategoryServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for plato.PlatoAnswerCategoryService.GetAnswerCategoryById.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getAnswerCategoryById(request: AnswerCategoryIdRequest):
        AnswerCategoryResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerCategoryService.GetAnswerCategoryById is unimplemented"))

    /**
     * Returns the response to an RPC for
     * plato.PlatoAnswerCategoryService.ListAnswerCategoriesByAnswerId.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listAnswerCategoriesByAnswerId(request: AnswerIdRequest):
        ListAnswerCategoriesResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerCategoryService.ListAnswerCategoriesByAnswerId is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerCategoryService.InsertAnswerCategory.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun insertAnswerCategory(request: InsertAnswerCategoryRequest):
        AnswerCategoryResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerCategoryService.InsertAnswerCategory is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerCategoryService.UpdateAnswerCategory.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateAnswerCategory(request: UpdateAnswerCategoryRequest):
        AnswerCategoryResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerCategoryService.UpdateAnswerCategory is unimplemented"))

    /**
     * Returns the response to an RPC for plato.PlatoAnswerCategoryService.DeleteAnswerCategory.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteAnswerCategory(request: AnswerCategoryIdRequest): GenericResponse
        = throw
        StatusException(UNIMPLEMENTED.withDescription("Method plato.PlatoAnswerCategoryService.DeleteAnswerCategory is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerCategoryServiceGrpc.getGetAnswerCategoryByIdMethod(),
      implementation = ::getAnswerCategoryById
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerCategoryServiceGrpc.getListAnswerCategoriesByAnswerIdMethod(),
      implementation = ::listAnswerCategoriesByAnswerId
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerCategoryServiceGrpc.getInsertAnswerCategoryMethod(),
      implementation = ::insertAnswerCategory
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerCategoryServiceGrpc.getUpdateAnswerCategoryMethod(),
      implementation = ::updateAnswerCategory
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = PlatoAnswerCategoryServiceGrpc.getDeleteAnswerCategoryMethod(),
      implementation = ::deleteAnswerCategory
    )).build()
  }
}
