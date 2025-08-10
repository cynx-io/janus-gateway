package com.philyra

import com.core.GenericResponse
import com.philyra.ResumeServiceGrpc.getServiceDescriptor
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
 * Holder for Kotlin coroutine-based client and server APIs for philyra.ResumeService.
 */
public object ResumeServiceGrpcKt {
  public const val SERVICE_NAME: String = ResumeServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val createResumeMethod: MethodDescriptor<CreateResumeRequest, ResumeResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getCreateResumeMethod()

  public val getResumeMethod: MethodDescriptor<GetResumeRequest, ResumeResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getGetResumeMethod()

  public val updateResumeMethod: MethodDescriptor<UpdateResumeRequest, ResumeResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getUpdateResumeMethod()

  public val listResumesMethod: MethodDescriptor<ListResumesRequest, ListResumesResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getListResumesMethod()

  public val deleteResumeMethod: MethodDescriptor<DeleteResumeRequest, GenericResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getDeleteResumeMethod()

  public val deleteResumeItemMethod: MethodDescriptor<DeleteResumeItemRequest, GenericResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getDeleteResumeItemMethod()

  public val getPersonalInfoMethod:
      MethodDescriptor<GetPersonalInfoRequest, GetPersonalInfoResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getGetPersonalInfoMethod()

  public val getExperienceMethod: MethodDescriptor<GetExperienceRequest, GetExperienceResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getGetExperienceMethod()

  public val getEducationMethod: MethodDescriptor<GetEducationRequest, GetEducationResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getGetEducationMethod()

  public val getSkillsMethod: MethodDescriptor<GetSkillsRequest, GetSkillsResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getGetSkillsMethod()

  public val upsertPersonalInfoMethod: MethodDescriptor<UpsertPersonalInfoRequest, GenericResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getUpsertPersonalInfoMethod()

  public val upsertExperienceMethod: MethodDescriptor<UpsertExperienceRequest, GenericResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getUpsertExperienceMethod()

  public val upsertEducationMethod: MethodDescriptor<UpsertEducationRequest, GenericResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getUpsertEducationMethod()

  public val upsertSkillsMethod: MethodDescriptor<UpsertSkillsRequest, GenericResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getUpsertSkillsMethod()

  public val generateResumeMethod: MethodDescriptor<GenerateResumeRequest, GenerateResumeResponse>
    @JvmStatic
    get() = ResumeServiceGrpc.getGenerateResumeMethod()

  /**
   * A stub for issuing RPCs to a(n) philyra.ResumeService service as suspending coroutines.
   */
  @StubFor(ResumeServiceGrpc::class)
  public class ResumeServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<ResumeServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions): ResumeServiceCoroutineStub =
        ResumeServiceCoroutineStub(channel, callOptions)

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
    public suspend fun createResume(request: CreateResumeRequest, headers: Metadata = Metadata()):
        ResumeResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getCreateResumeMethod(),
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
    public suspend fun getResume(request: GetResumeRequest, headers: Metadata = Metadata()):
        ResumeResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getGetResumeMethod(),
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
    public suspend fun updateResume(request: UpdateResumeRequest, headers: Metadata = Metadata()):
        ResumeResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getUpdateResumeMethod(),
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
    public suspend fun listResumes(request: ListResumesRequest, headers: Metadata = Metadata()):
        ListResumesResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getListResumesMethod(),
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
    public suspend fun deleteResume(request: DeleteResumeRequest, headers: Metadata = Metadata()):
        GenericResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getDeleteResumeMethod(),
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
    public suspend fun deleteResumeItem(request: DeleteResumeItemRequest, headers: Metadata =
        Metadata()): GenericResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getDeleteResumeItemMethod(),
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
    public suspend fun getPersonalInfo(request: GetPersonalInfoRequest, headers: Metadata =
        Metadata()): GetPersonalInfoResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getGetPersonalInfoMethod(),
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
    public suspend fun getExperience(request: GetExperienceRequest, headers: Metadata = Metadata()):
        GetExperienceResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getGetExperienceMethod(),
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
    public suspend fun getEducation(request: GetEducationRequest, headers: Metadata = Metadata()):
        GetEducationResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getGetEducationMethod(),
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
    public suspend fun getSkills(request: GetSkillsRequest, headers: Metadata = Metadata()):
        GetSkillsResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getGetSkillsMethod(),
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
    public suspend fun upsertPersonalInfo(request: UpsertPersonalInfoRequest, headers: Metadata =
        Metadata()): GenericResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getUpsertPersonalInfoMethod(),
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
    public suspend fun upsertExperience(request: UpsertExperienceRequest, headers: Metadata =
        Metadata()): GenericResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getUpsertExperienceMethod(),
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
    public suspend fun upsertEducation(request: UpsertEducationRequest, headers: Metadata =
        Metadata()): GenericResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getUpsertEducationMethod(),
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
    public suspend fun upsertSkills(request: UpsertSkillsRequest, headers: Metadata = Metadata()):
        GenericResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getUpsertSkillsMethod(),
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
    public suspend fun generateResume(request: GenerateResumeRequest, headers: Metadata =
        Metadata()): GenerateResumeResponse = unaryRpc(
      channel,
      ResumeServiceGrpc.getGenerateResumeMethod(),
      request,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the philyra.ResumeService service based on Kotlin coroutines.
   */
  public abstract class ResumeServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns the response to an RPC for philyra.ResumeService.CreateResume.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun createResume(request: CreateResumeRequest): ResumeResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.CreateResume is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.GetResume.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getResume(request: GetResumeRequest): ResumeResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.GetResume is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.UpdateResume.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun updateResume(request: UpdateResumeRequest): ResumeResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.UpdateResume is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.ListResumes.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun listResumes(request: ListResumesRequest): ListResumesResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.ListResumes is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.DeleteResume.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteResume(request: DeleteResumeRequest): GenericResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.DeleteResume is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.DeleteResumeItem.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun deleteResumeItem(request: DeleteResumeItemRequest): GenericResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.DeleteResumeItem is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.GetPersonalInfo.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getPersonalInfo(request: GetPersonalInfoRequest):
        GetPersonalInfoResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.GetPersonalInfo is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.GetExperience.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getExperience(request: GetExperienceRequest): GetExperienceResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.GetExperience is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.GetEducation.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getEducation(request: GetEducationRequest): GetEducationResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.GetEducation is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.GetSkills.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun getSkills(request: GetSkillsRequest): GetSkillsResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.GetSkills is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.UpsertPersonalInfo.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun upsertPersonalInfo(request: UpsertPersonalInfoRequest): GenericResponse
        = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.UpsertPersonalInfo is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.UpsertExperience.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun upsertExperience(request: UpsertExperienceRequest): GenericResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.UpsertExperience is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.UpsertEducation.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun upsertEducation(request: UpsertEducationRequest): GenericResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.UpsertEducation is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.UpsertSkills.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun upsertSkills(request: UpsertSkillsRequest): GenericResponse = throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.UpsertSkills is unimplemented"))

    /**
     * Returns the response to an RPC for philyra.ResumeService.GenerateResume.
     *
     * If this method fails with a [StatusException], the RPC will fail with the corresponding
     * [io.grpc.Status].  If this method fails with a [java.util.concurrent.CancellationException],
     * the RPC will fail
     * with status `Status.CANCELLED`.  If this method fails for any other reason, the RPC will
     * fail with `Status.UNKNOWN` with the exception as a cause.
     *
     * @param request The request from the client.
     */
    public open suspend fun generateResume(request: GenerateResumeRequest): GenerateResumeResponse =
        throw
        StatusException(UNIMPLEMENTED.withDescription("Method philyra.ResumeService.GenerateResume is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getCreateResumeMethod(),
      implementation = ::createResume
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getGetResumeMethod(),
      implementation = ::getResume
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getUpdateResumeMethod(),
      implementation = ::updateResume
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getListResumesMethod(),
      implementation = ::listResumes
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getDeleteResumeMethod(),
      implementation = ::deleteResume
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getDeleteResumeItemMethod(),
      implementation = ::deleteResumeItem
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getGetPersonalInfoMethod(),
      implementation = ::getPersonalInfo
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getGetExperienceMethod(),
      implementation = ::getExperience
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getGetEducationMethod(),
      implementation = ::getEducation
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getGetSkillsMethod(),
      implementation = ::getSkills
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getUpsertPersonalInfoMethod(),
      implementation = ::upsertPersonalInfo
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getUpsertExperienceMethod(),
      implementation = ::upsertExperience
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getUpsertEducationMethod(),
      implementation = ::upsertEducation
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getUpsertSkillsMethod(),
      implementation = ::upsertSkills
    ))
      .addMethod(unaryServerMethodDefinition(
      context = this.context,
      descriptor = ResumeServiceGrpc.getGenerateResumeMethod(),
      implementation = ::generateResume
    )).build()
  }
}
