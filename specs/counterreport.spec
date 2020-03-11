# Model
model:
  rest_name: counterreport
  resource_name: counterreports
  entity_name: CounterReport
  package: zack
  group: core/enforcer
  description: Post a new counter tracing report.

# Attributes
attributes:
  v1:
  - name: AckInUnknownState
    description: Counter for sending finack ack received in unknown connection state.
    type: integer
    exposed: true
    default_value: 0

  - name: AckInvalidFormat
    description: Counter for ack packet dropped because of invalid format.
    type: integer
    exposed: true
    default_value: 0

  - name: AckRejected
    description: Counter for reject ack packet as per policy.
    type: integer
    exposed: true
    default_value: 0

  - name: AckSigValidationFailed
    description: Counter for ack packet dropped because signature validation failed.
    type: integer
    exposed: true
    default_value: 0

  - name: AckTCPNoTCPAuthOption
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: ConnectionsProcessed
    description: Counter for connections processed.
    type: integer
    exposed: true
    default_value: 0

  - name: ContextIDNotFound
    description: Counter for unable to find ContextID.
    type: integer
    exposed: true
    default_value: 0

  - name: DroppedExternalService
    description: |-
      Counter for no acls found for external services. dropping application syn
      packet.
    type: integer
    exposed: true
    default_value: 0

  - name: InvalidConnState
    description: Counter for invalid connection state.
    type: integer
    exposed: true
    default_value: 0

  - name: InvalidNetState
    description: Counter for invalid net state.
    type: integer
    exposed: true
    default_value: 0

  - name: InvalidProtocol
    description: Counter for invalid protocol.
    type: integer
    exposed: true
    default_value: 0

  - name: InvalidSynAck
    description: Counter for pu is already dead - drop synack packet.
    type: integer
    exposed: true
    default_value: 0

  - name: MarkNotFound
    description: Counter for pu mark not found.
    type: integer
    exposed: true
    default_value: 0

  - name: NetSynNotSeen
    description: Counter for network syn packet was not seen.
    type: integer
    exposed: true
    default_value: 0

  - name: NoConnFound
    description: Counter for no context or connection found.
    type: integer
    exposed: true
    default_value: 0

  - name: NonPUTraffic
    description: Counter for traffic that belongs to a non PU process.
    type: integer
    exposed: true
    default_value: 0

  - name: OutOfOrderSynAck
    description: Counter for synack for flow with processed finack.
    type: integer
    exposed: true
    default_value: 0

  - name: PortNotFound
    description: Counter for port not found.
    type: integer
    exposed: true
    default_value: 0

  - name: RejectPacket
    description: Counter for reject the packet as per policy.
    type: integer
    exposed: true
    default_value: 0

  - name: ServicePostprocessorFailed
    description: Counter for post service processing failed for network packet.
    type: integer
    exposed: true
    default_value: 0

  - name: ServicePreprocessorFailed
    description: Counter for network packets that failed preprocessing.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckBadClaims
    description: Counter for synack packet dropped because of bad claims.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckClaimsMisMatch
    description: Counter for synack packet dropped because of encryption mismatch.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckDroppedExternalService
    description: Counter for synack from external service dropped.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckInvalidFormat
    description: Counter for synack packet dropped because of invalid format.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckMissingClaims
    description: Counter for synack packet dropped because of no claims.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckMissingToken
    description: Counter for synack packet dropped because of missing token.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckNoTCPAuthOption
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: SynAckRejected
    description: Counter for dropping because of reject rule on transmitter.
    type: integer
    exposed: true
    default_value: 0

  - name: SynDroppedInvalidFormat
    description: Counter for syn packet dropped because of invalid format.
    type: integer
    exposed: true
    default_value: 0

  - name: SynDroppedInvalidToken
    description: Counter for syn packet dropped because of invalid token.
    type: integer
    exposed: true
    default_value: 0

  - name: SynDroppedNoClaims
    description: Counter for syn packet dropped because of no claims.
    type: integer
    exposed: true
    default_value: 0

  - name: SynDroppedTCPOption
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: SynRejectPacket
    description: Counter for syn dropped due to policy.
    type: integer
    exposed: true
    default_value: 0

  - name: SynUnexpectedPacket
    description: Counter for received syn packet from unknown pu.
    type: integer
    exposed: true
    default_value: 0

  - name: TCPAuthNotFound
    description: Counter for tcp authentication option not found.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPAckInvalidSignature
    description: Counter for dropped udp ack invalid signature.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPConnectionsProcessed
    description: Counter for number of processed UDP connections.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropContextNotFound
    description: Counter for dropped UDP data packets with no context.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropFin
    description: Counter for dropped udp FIN handshake packets.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropInNfQueue
    description: Counter for dropped UDP in NfQueue.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropNoConnection
    description: Counter for dropped UDP data packets with no connection.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropPacket
    description: Counter for dropped UDP data packets.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropQueueFull
    description: Counter for dropped UDP Queue Full.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPDropSynAck
    description: Counter for dropped udp synack handshake packets.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPInvalidNetState
    description: Counter for udp packets received in invalid network state.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPPostProcessingFailed
    description: Counter for UDP packets failing postprocessing.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPPreProcessingFailed
    description: Counter for UDP packets failing preprocessing.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPRejected
    description: Counter for UDP packets dropped due to policy.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynAckDropBadClaims
    description: Counter for dropped udp synack bad claims.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynAckMissingClaims
    description: Counter for dropped udp synack missing claims.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynAckPolicy
    description: Counter for dropped udp synack bad claims.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynDrop
    description: Counter for dropped udp syn transmits.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynDropPolicy
    description: Counter for dropped udp syn policy.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynInvalidToken
    description: Counter for dropped udp FIN handshake packets.
    type: integer
    exposed: true
    default_value: 0

  - name: UDPSynMissingClaims
    description: Counter for dropped UDP SYN missing claims.
    type: integer
    exposed: true
    default_value: 0

  - name: UnknownError
    description: Counter for unknown error.
    type: integer
    exposed: true
    default_value: 0

  - name: enforcerID
    description: Identifier of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxx-xxx-xxxx

  - name: enforcerNamespace
    description: Namespace of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace

  - name: processingUnitID
    description: PUID is the ID of the PU reporting the counter.
    type: string
    exposed: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: processingUnitNamespace
    description: Namespace of the PU reporting the counter.
    type: string
    exposed: true
    example_value: /my/namespace
    filterable: true

  - name: timestamp
    description: Timestamp is the date of the report.
    type: time
    exposed: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: DuplicateAckDrop
    description: Counter for duplicate ack drop.
    type: integer
    exposed: true
    default_value: 0
