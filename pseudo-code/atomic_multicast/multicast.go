// pseudo-code
type message struct{
	*mq.message 
	ACKMessage map[int]*mq.message
	NoACKMessage map[int]*node
	MessageId int
	SequenceNumber int
}
func SendAtomicMulticase(){
	//Fetch next messaage as M from MQ
	if M!=nil{
		M.SequenceNumber=LastestSequenceNumber+1
	//Initialize M for Multicast
	for (M.Limit_time!=config.time_out){
			for _,N :=range Node_group{
				IsSendMessage(M,N)
			}	
	}
	delete(M)
}
func SendMessage(M message,N node)bool{
	//INITIALIZE
	if err:=Verify_ACK_Node(N);err!=nil{
		M.SendStatus=NOT_SEND
		M.Description=err
		DisplayERR(err)
		return false
	}
	// GET SOCKET DESCRIPTOR D for sending message
	if D==nil{
		DisplayERR(err)
		M.SendStatus=NOT_SEND
		return false
	}
	//SEND M TO N USING SOCKET 
	if !MQSendACK(m message){
		M.SendStatus=NOT_SEND
		M.SendStatus=NO_ACK
		return false 
	}
	M.SendStatus=ALREADY_SEND
	close(D)
	Realse(date)
	return true 
}
func Varify_ACK_Node(N node)error  {
	
}
func IsMQSendACK(m messaage)bool{

}
func Realse() {
	
}

func HandleACkMessage(AckMessage message)  {
	if AckMessage.received{
		return
	}
	retransmission(AckMessage)
	RecordACKMessage(AckMessage)
	IncreaseACKMessaage()
	retransmission(AckMessage)
	
}
func DetermineFinalSeqenceNumber()int{
	//Initialize Max SeqenceNumber
	for (ACKMessage M IN ACKMessage map){
		if (MaxSequenceNumber<ACKMessage.SequenceNumber){
			MaxSequenceNumber=ACKMessage.SequenceNumber
		}
	}
	return MaxSequenceNumber
}
func SendFinalSequenceNumber(FinalSequenceNumber int){
	//initialize 
	for (N node in group view){
		SendFinalSequenceNumber(FinalSequenceNumber,N)
	}
	Realse(data)
}
func HandleDateMessage(DataMessage M,UUID int)  {
	PutIntoUndeliverableQueue(M,UUID)
	
}
func PutIntoUndeliverableQueue(DataMessage M,UUID int){
	//Map LastestSequenceNumber TO M.sender
	//Mark UUID Of M as undeliverable
}
func ReplyACKMessage (DataMessage M){
	//Build new ACK Message of M as Message 
	if Msg!=nil{
		//send Msg to M.sender
	}
	Realse(data)
}
func HandlerFinalSequenceMessage( M Message){
	//use M->UUID TO FIND CORRESPONDING MSG
	//REMOVE OLD SEQUENCE NUMBER OF MESSSAGE
	//ASSGIN M.FinalSequenceNumber to message 
	//Deliver Msg to application
}
func DeliverToApplication(M message)  {
	//for(M in deliverable queue && M.status id deliverable){
	//push m into message queue
	//Erase M frome current deliverable queue
	}
	
}
func retransmission(M message){

}
func RecordACKMessage(M message){

}
func IncreaseACKMessaage()  {
	
}
func StoreFailSendMessage(M message,n node)  {
	
}