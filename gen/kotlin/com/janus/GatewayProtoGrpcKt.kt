package com.janus

import com.janus.GatewayServiceGrpc.getServiceDescriptor
import io.grpc.CallOptions
import io.grpc.CallOptions.DEFAULT
import io.grpc.Channel
import io.grpc.ServerServiceDefinition
import io.grpc.ServerServiceDefinition.builder
import io.grpc.ServiceDescriptor
import io.grpc.kotlin.AbstractCoroutineServerImpl
import io.grpc.kotlin.AbstractCoroutineStub
import io.grpc.kotlin.StubFor
import kotlin.String
import kotlin.coroutines.CoroutineContext
import kotlin.coroutines.EmptyCoroutineContext
import kotlin.jvm.JvmOverloads
import kotlin.jvm.JvmStatic

/**
 * Holder for Kotlin coroutine-based client and server APIs for janus.GatewayService.
 */
public object GatewayServiceGrpcKt {
  public const val SERVICE_NAME: String = GatewayServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  /**
   * A stub for issuing RPCs to a(n) janus.GatewayService service as suspending coroutines.
   */
  @StubFor(GatewayServiceGrpc::class)
  public class GatewayServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<GatewayServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): GatewayServiceCoroutineStub =
        GatewayServiceCoroutineStub(channel, callOptions)
  }

  /**
   * Skeletal implementation of the janus.GatewayService service based on Kotlin coroutines.
   */
  public abstract class GatewayServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    final override fun bindService(): ServerServiceDefinition =
        builder(getServiceDescriptor()).build()
  }
}
