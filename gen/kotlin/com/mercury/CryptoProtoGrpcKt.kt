package com.mercury

import com.mercury.MercuryCryptoServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for mercury.MercuryCryptoService.
 */
public object MercuryCryptoServiceGrpcKt {
  public const val SERVICE_NAME: String = MercuryCryptoServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val searchCoinMethod: MethodDescriptor<SearchCoinRequest, SearchCoinResponse>
    @JvmStatic
    get() = MercuryCryptoServiceGrpc.getSearchCoinMethod()

  public val getCoinRiskMethod: MethodDescriptor<GetCoinRiskRequest, GetCoinRiskResponse>
    @JvmStatic
    get() = MercuryCryptoServiceGrpc.getGetCoinRiskMethod()

  /**
   * A stub for issuing RPCs to a(n) mercury.MercuryCryptoService service as suspending coroutines.
   */
  @StubFor(MercuryCryptoServiceGrpc::class)
  public class MercuryCryptoServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<MercuryCryptoServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions):
        MercuryCryptoServiceCoroutineStub = MercuryCryptoServiceCoroutineStub(channel, callOptions)

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
    public suspend fun searchCoin(request: SearchCoinRequest, headers: Metadata = Metadata()):
        SearchCoinResponse = unaryRpc(
      channel,
      MercuryCryptoServiceGrpc.getSearchCoinMethod(),
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
    public suspend fun getCoinRisk(request: GetCoinRiskRequest, headers: Metadata = Metadata()):
        GetCoinRiskResponse = unaryRpc(
      channel,
      MercuryCryptoServiceGrpc.getGetCoinRiskMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the mercury.MercuryCryptoService service based on Kotlin coroutines.
   */
  public abstract class MercuryCryptoServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for mercury.MercuryCryptoService.SearchCoin.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun searchCoin(request: SearchCoinRequest): SearchCoinResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method mercury.MercuryCryptoService.SearchCoin is unimplemented"))

    /**
     * Returns the response to an RPC for mercury.MercuryCryptoService.GetCoinRisk.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getCoinRisk(request: GetCoinRiskRequest): GetCoinRiskResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method mercury.MercuryCryptoService.GetCoinRisk is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = MercuryCryptoServiceGrpc.getSearchCoinMethod(),
      implementation = ::searchCoin
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = MercuryCryptoServiceGrpc.getGetCoinRiskMethod(),
      implementation = ::getCoinRisk
    )).build()
  }
}
