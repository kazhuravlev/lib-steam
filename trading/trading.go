package trading

type TradeOfferState uint8

const (
	// TradeOfferStateInvalid - Invalid
	// Неправильный. непонятно что означает
	TradeOfferStateInvalid TradeOfferState = 1
	// TradeOfferStateActive - This trade offer has been sent, neither party has acted on it yet.
	// Активный. значит, преложение еще не принято. имеется класс в DOM
	TradeOfferStateActive TradeOfferState = 2
	// TradeOfferStateAccepted - The trade offer was accepted by the recipient and items were exchanged.
	// Предложение принято. имеется класс в DOM
	TradeOfferStateAccepted TradeOfferState = 3
	// TradeOfferStateCountered - The recipient made a counter offer
	TradeOfferStateCountered TradeOfferState = 4
	// TradeOfferStateExpired - The trade offer was not accepted before the expiration date
	// Протух. возможно по окончанию expiration_time сделка превращается в тыкву
	TradeOfferStateExpired TradeOfferState = 5
	// TradeOfferStateCanceled - The sender cancelled the offer
	// Отменен. Отмена производится со стороны отправителя. Он как бы отменяет то, что сам создал
	TradeOfferStateCanceled TradeOfferState = 6
	// TradeOfferStateDeclined - The recipient declined the offer
	// Отклонен. отклоняется чухой входящий оффер.
	TradeOfferStateDeclined TradeOfferState = 7
	// TradeOfferStateInvalidItems - Some of the items in the offer are no longer available (indicated by the missing flag in the
	// output)
	// Некорректные предметы. хз что такое. Вохможно, такой стаус сделка получает если
	TradeOfferStateInvalidItems TradeOfferState = 8
	// TradeOfferStateCreatedNeedsConfirmation - The offer hasn't been sent yet and is awaiting email/mobile confirmation. The
	// offer is only visible to the sender.
	TradeOfferStateCreatedNeedsConfirmation TradeOfferState = 9
	// TradeOfferStateCanceledBySecondFactor - Either party canceled the offer via email/mobile. The offer is visible to both
	// parties, even if the sender canceled it before it was sent.
	TradeOfferStateCanceledBySecondFactor TradeOfferState = 10
	// TradeOfferStateInEscrow - The trade has been placed on hold. The items involved in the trade have all been removed from
	// both parties' inventories and will be automatically delivered in the future.
	TradeOfferStateInEscrow TradeOfferState = 11
)

var (
	// TradeOfferStates список всех доступных состояний сделки
	TradeOfferStates = []TradeOfferState{
		TradeOfferStateInvalid,
		TradeOfferStateActive,
		TradeOfferStateAccepted,
		TradeOfferStateCountered,
		TradeOfferStateExpired,
		TradeOfferStateCanceled,
		TradeOfferStateDeclined,
		TradeOfferStateInvalidItems,
		TradeOfferStateCreatedNeedsConfirmation,
		TradeOfferStateCanceledBySecondFactor,
		TradeOfferStateInEscrow,
	}
)

type ConfirmationMethod uint

const (
	// ConfirmationMethod_Invalid - Invalid
	ConfirmationMethodInvalid ConfirmationMethod = 0
	// ConfirmationMethod_Email - An email was sent with details on how to confirm the trade offer
	ConfirmationMethodEmail ConfirmationMethod = 1
	// ConfirmationMethod_MobileApp - The trade offer may be confirmed via the mobile app
	ConfirmationMethodMobileApp ConfirmationMethod = 2
)
