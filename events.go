package zstripe

// List of webhook events; API version 2020-08-27.
const (
	// Account status or property has changed.
	EventAccountUpdated = "account.updated"

	// User authorizes an application. Sent to the related application only.
	EventAccountApplicationAuthorized = "account.application.authorized"

	// User deauthorizes an application. Sent to the related application only.
	EventAccountApplicationDeauthorized = "account.application.deauthorized"

	// External account is created.
	EventAccountExternalAccountCreated = "account.external_account.created"

	// External account is deleted.
	EventAccountExternalAccountDeleted = "account.external_account.deleted"

	// External account is updated.
	EventAccountExternalAccountUpdated = "account.external_account.updated"

	// Application fee is created on a charge.
	EventApplicationFeeCreated = "application_fee.created"

	// Application fee is refunded, whether from refunding a charge or from
	// refunding the application fee directly. This includes partial refunds.
	EventApplicationFeeRefunded = "application_fee.refunded"

	// Application fee refund is updated.
	EventApplicationFeeRefundUpdated = "application_fee.refund.updated"

	// Your Stripe balance has been updated (e.g., when a charge is available to
	// be paid out). By default, Stripe automatically transfers funds in your
	// balance to your bank account on a daily basis.
	EventBalanceAvailable = "balance.available"

	// Portal configuration is created.
	EventBillingPortalConfigurationCreated = "billing_portal.configuration.created"

	// Portal configuration is updated.
	EventBillingPortalConfigurationUpdated = "billing_portal.configuration.updated"

	// Capability has new requirements or a new status.
	EventCapabilityUpdated = "capability.updated"

	// Previously uncaptured charge is captured.
	EventChargeCaptured = "charge.captured"

	// Uncaptured charge expires.
	EventChargeExpired = "charge.expired"

	// Failed charge attempt occurs.
	EventChargeFailed = "charge.failed"

	// Pending charge is created.
	EventChargePending = "charge.pending"

	// Charge is refunded, including partial refunds.
	EventChargeRefunded = "charge.refunded"

	// New charge is created and is successful.
	EventChargeSucceeded = "charge.succeeded"

	// Charge description or metadata is updated.
	EventChargeUpdated = "charge.updated"

	// Dispute is closed and the dispute status changes to lost, warning_closed,
	// or won.
	EventChargeDisputeClosed = "charge.dispute.closed"

	// Customer disputes a charge with their bank.
	EventChargeDisputeCreated = "charge.dispute.created"

	// Funds are reinstated to your account after a dispute is closed. This
	// includes partially refunded payments.
	EventChargeDisputeFundsReinstated = "charge.dispute.funds_reinstated"

	// Funds are removed from your account due to a dispute.
	EventChargeDisputeFundsWithdrawn = "charge.dispute.funds_withdrawn"

	// Dispute is updated (usually with evidence).
	EventChargeDisputeUpdated = "charge.dispute.updated"

	// Refund is updated, on selected payment methods.
	EventChargeRefundUpdated = "charge.refund.updated"

	// Payment intent using a delayed payment method fails.
	EventCheckoutSessionAsyncPaymentFailed = "checkout.session.async_payment_failed"

	// Payment intent using a delayed payment method finally succeeds.
	EventCheckoutSessionAsyncPaymentSucceeded = "checkout.session.async_payment_succeeded"

	// Checkout Session has been successfully completed.
	EventCheckoutSessionCompleted = "checkout.session.completed"

	// Coupon is created.
	EventCouponCreated = "coupon.created"

	// Coupon is deleted.
	EventCouponDeleted = "coupon.deleted"

	// Coupon is updated.
	EventCouponUpdated = "coupon.updated"

	// Credit note is created.
	EventCreditNoteCreated = "credit_note.created"

	// Credit note is updated.
	EventCreditNoteUpdated = "credit_note.updated"

	// Credit note is voided.
	EventCreditNoteVoided = "credit_note.voided"

	// New customer is created.
	EventCustomerCreated = "customer.created"

	// Customer is deleted.
	EventCustomerDeleted = "customer.deleted"

	// Any property of a customer changes.
	EventCustomerUpdated = "customer.updated"

	// Coupon is attached to a customer.
	EventCustomerDiscountCreated = "customer.discount.created"

	// Coupon is removed from a customer.
	EventCustomerDiscountDeleted = "customer.discount.deleted"

	// Customer is switched from one coupon to another.
	EventCustomerDiscountUpdated = "customer.discount.updated"

	// New source is created for a customer.
	EventCustomerSourceCreated = "customer.source.created"

	// Source is removed from a customer.
	EventCustomerSourceDeleted = "customer.source.deleted"

	// Card or source will expire at the end of the month.
	EventCustomerSourceExpiring = "customer.source.expiring"

	// Source's details are changed.
	EventCustomerSourceUpdated = "customer.source.updated"

	// Customer is signed up for a new plan.
	EventCustomerSubscriptionCreated = "customer.subscription.created"

	// Customer's subscription ends.
	EventCustomerSubscriptionDeleted = "customer.subscription.deleted"

	// Customer's subscription's pending update is applied, and the subscription is updated.
	EventCustomerSubscriptionPendingUpdateApplied = "customer.subscription.pending_update_applied"

	// Customer's subscription's pending update expires before the related invoice is paid.
	EventCustomerSubscriptionPendingUpdateExpired = "customer.subscription.pending_update_expired"

	// Three days before a subscription's trial period is scheduled to end, or
	// when a trial is ended immediately (using trial_end=now).
	EventCustomerSubscriptionTrialWillEnd = "customer.subscription.trial_will_end"

	// Subscription changes (e.g., switching from one plan to another, or
	// changing the status from trial to active).
	EventCustomerSubscriptionUpdated = "customer.subscription.updated"

	// Tax ID is created for a customer.
	EventCustomerTaxIdCreated = "customer.tax_id.created"

	// Tax ID is deleted from a customer.
	EventCustomerTaxIdDeleted = "customer.tax_id.deleted"

	// Customer's tax ID is updated.
	EventCustomerTaxIdUpdated = "customer.tax_id.updated"

	// New Stripe-generated file is available for your account.
	EventFileCreated = "file.created"

	// New invoice is created. To learn how webhooks can be used with this
	// event, and how they can affect it, see Using Webhooks with Subscriptions.
	EventInvoiceCreated = "invoice.created"

	// Draft invoice is deleted.
	EventInvoiceDeleted = "invoice.deleted"

	// Draft invoice cannot be finalized. See the invoice’s last finalization error for details.
	EventInvoiceFinalizationFailed = "invoice.finalization_failed"

	// Draft invoice is finalized and updated to be an open invoice.
	EventInvoiceFinalized = "invoice.finalized"

	// Invoice is marked uncollectible.
	EventInvoiceMarkedUncollectible = "invoice.marked_uncollectible"

	// Invoice payment attempt succeeds or an invoice is marked as paid out-of-band.
	EventInvoicePaid = "invoice.paid"

	// Invoice payment attempt requires further user action to complete.
	EventInvoicePaymentActionRequired = "invoice.payment_action_required"

	// Invoice payment attempt fails, due either to a declined payment or to the lack of a stored payment method.
	EventInvoicePaymentFailed = "invoice.payment_failed"

	// Invoice payment attempt succeeds.
	EventInvoicePaymentSucceeded = "invoice.payment_succeeded"

	// Invoice email is sent out.
	EventInvoiceSent = "invoice.sent"

	// X number of days before a subscription is scheduled to create an invoice
	// that is automatically charged&mdash;where X is determined by your
	// subscriptions settings. Note: The received Invoice object will not have
	// an invoice ID.
	EventInvoiceUpcoming = "invoice.upcoming"

	// Invoice changes (e.g., the invoice amount).
	EventInvoiceUpdated = "invoice.updated"

	// Invoice is voided.
	EventInvoiceVoided = "invoice.voided"

	// Invoice item is created.
	EventInvoiceitemCreated = "invoiceitem.created"

	// Invoice item is deleted.
	EventInvoiceitemDeleted = "invoiceitem.deleted"

	// Invoice item is updated.
	EventInvoiceitemUpdated = "invoiceitem.updated"

	// Authorization is created.
	EventIssuingAuthorizationCreated = "issuing_authorization.created"

	// Represents a synchronous request for authorization, see Using your
	// integration to handle authorization requests.
	EventIssuingAuthorizationRequest = "issuing_authorization.request"

	// Authorization is updated.
	EventIssuingAuthorizationUpdated = "issuing_authorization.updated"

	// Card is created.
	EventIssuingCardCreated = "issuing_card.created"

	// Card is updated.
	EventIssuingCardUpdated = "issuing_card.updated"

	// Cardholder is created.
	EventIssuingCardholderCreated = "issuing_cardholder.created"

	// Cardholder is updated.
	EventIssuingCardholderUpdated = "issuing_cardholder.updated"

	// Dispute is won, lost or expired.
	EventIssuingDisputeClosed = "issuing_dispute.closed"

	// Dispute is created.
	EventIssuingDisputeCreated = "issuing_dispute.created"

	// Funds are reinstated to your account for an Issuing dispute.
	EventIssuingDisputeFundsReinstated = "issuing_dispute.funds_reinstated"

	// Dispute is submitted.
	EventIssuingDisputeSubmitted = "issuing_dispute.submitted"

	// Dispute is updated.
	EventIssuingDisputeUpdated = "issuing_dispute.updated"

	// Issuing transaction is created.
	EventIssuingTransactionCreated = "issuing_transaction.created"

	// Issuing transaction is updated.
	EventIssuingTransactionUpdated = "issuing_transaction.updated"

	// Mandate is updated.
	EventMandateUpdated = "mandate.updated"

	// Order is created.
	EventOrderCreated = "order.created"

	// Order payment attempt fails.
	EventOrderPaymentFailed = "order.payment_failed"

	// Order payment attempt succeeds.
	EventOrderPaymentSucceeded = "order.payment_succeeded"

	// Order is updated.
	EventOrderUpdated = "order.updated"

	// Order return is created.
	EventOrderReturnCreated = "order_return.created"

	// PaymentIntent has funds to be captured. Check the amount_capturable
	// property on the PaymentIntent to determine the amount that can be
	// captured. You may capture the PaymentIntent with an amount_to_capture
	// value up to the specified amount. Learn more about capturing
	// PaymentIntents.
	EventPaymentIntentAmountCapturableUpdated = "payment_intent.amount_capturable_updated"

	// PaymentIntent is canceled.
	EventPaymentIntentCanceled = "payment_intent.canceled"

	// New PaymentIntent is created.
	EventPaymentIntentCreated = "payment_intent.created"

	// PaymentIntent has failed the attempt to create a payment method or a payment.
	EventPaymentIntentPaymentFailed = "payment_intent.payment_failed"

	// PaymentIntent has started processing.
	EventPaymentIntentProcessing = "payment_intent.processing"

	// PaymentIntent transitions to requires_action state
	EventPaymentIntentRequiresAction = "payment_intent.requires_action"

	// PaymentIntent has successfully completed payment.
	EventPaymentIntentSucceeded = "payment_intent.succeeded"

	// New payment method is attached to a customer.
	EventPaymentMethodAttached = "payment_method.attached"

	// Payment method's details are automatically updated by the network.
	EventPaymentMethodAutomaticallyUpdated = "payment_method.automatically_updated"

	// Payment method is detached from a customer.
	EventPaymentMethodDetached = "payment_method.detached"

	// Payment method is updated via the PaymentMethod update API.
	EventPaymentMethodUpdated = "payment_method.updated"

	// Payout is canceled.
	EventPayoutCanceled = "payout.canceled"

	// Payout is created.
	EventPayoutCreated = "payout.created"

	// Payout attempt fails.
	EventPayoutFailed = "payout.failed"

	// Payout is expected to be available in the destination account. If the
	// payout fails, a payout.failed notification is also sent, at a later time.
	EventPayoutPaid = "payout.paid"

	// Payout is updated.
	EventPayoutUpdated = "payout.updated"

	// Person associated with an account is created.
	EventPersonCreated = "person.created"

	// Person associated with an account is deleted.
	EventPersonDeleted = "person.deleted"

	// Person associated with an account is updated.
	EventPersonUpdated = "person.updated"

	// Plan is created.
	EventPlanCreated = "plan.created"

	// Plan is deleted.
	EventPlanDeleted = "plan.deleted"

	// Plan is updated.
	EventPlanUpdated = "plan.updated"

	// Price is created.
	EventPriceCreated = "price.created"

	// Price is deleted.
	EventPriceDeleted = "price.deleted"

	// Price is updated.
	EventPriceUpdated = "price.updated"

	// Product is created.
	EventProductCreated = "product.created"

	// Product is deleted.
	EventProductDeleted = "product.deleted"

	// Product is updated.
	EventProductUpdated = "product.updated"

	// Promotion code is created.
	EventPromotionCodeCreated = "promotion_code.created"

	// Promotion code is updated.
	EventPromotionCodeUpdated = "promotion_code.updated"

	// Early fraud warning is created.
	EventRadarEarlyFraudWarningCreated = "radar.early_fraud_warning.created"

	// Early fraud warning is updated.
	EventRadarEarlyFraudWarningUpdated = "radar.early_fraud_warning.updated"

	// Recipient is created.
	EventRecipientCreated = "recipient.created"

	// Recipient is deleted.
	EventRecipientDeleted = "recipient.deleted"

	// Recipient is updated.
	EventRecipientUpdated = "recipient.updated"

	// Requested ReportRun failed to complete.
	EventReportingReportRunFailed = "reporting.report_run.failed"

	// Requested ReportRun completed succesfully.
	EventReportingReportRunSucceeded = "reporting.report_run.succeeded"

	// ReportType is updated (typically to indicate that a new day's data has come available).
	EventReportingReportTypeUpdated = "reporting.report_type.updated"

	// Review is closed. The review's reason field indicates why: approved,
	// disputed, refunded, or refunded_as_fraud.
	EventReviewClosed = "review.closed"

	// Review is opened.
	EventReviewOpened = "review.opened"

	// SetupIntent is canceled.
	EventSetupIntentCanceled = "setup_intent.canceled"

	// New SetupIntent is created.
	EventSetupIntentCreated = "setup_intent.created"

	// SetupIntent is in requires_action state.
	EventSetupIntentRequiresAction = "setup_intent.requires_action"

	// SetupIntent has failed the attempt to setup a payment method.
	EventSetupIntentSetupFailed = "setup_intent.setup_failed"

	// SetupIntent has successfully setup a payment method.
	EventSetupIntentSucceeded = "setup_intent.succeeded"

	// Sigma scheduled query run finishes.
	EventSigmaScheduledQueryRunCreated = "sigma.scheduled_query_run.created"

	// SKU is created.
	EventSkuCreated = "sku.created"

	// SKU is deleted.
	EventSkuDeleted = "sku.deleted"

	// SKU is updated.
	EventSkuUpdated = "sku.updated"

	// Source is canceled.
	EventSourceCanceled = "source.canceled"

	// Source transitions to chargeable.
	EventSourceChargeable = "source.chargeable"

	// Source fails.
	EventSourceFailed = "source.failed"

	// Source mandate notification method is set to manual.
	EventSourceMandateNotification = "source.mandate_notification"

	// Refund attributes are required on a receiver source to process a refund or a mispayment.
	EventSourceRefundAttributesRequired = "source.refund_attributes_required"

	// Source transaction is created.
	EventSourceTransactionCreated = "source.transaction.created"

	// Source transaction is updated.
	EventSourceTransactionUpdated = "source.transaction.updated"

	// Subscription schedule is canceled due to the underlying subscription
	// being canceled because of delinquency.
	EventSubscriptionScheduleAborted = "subscription_schedule.aborted"

	// Subscription schedule is canceled.
	EventSubscriptionScheduleCanceled = "subscription_schedule.canceled"

	// New subscription schedule is completed.
	EventSubscriptionScheduleCompleted = "subscription_schedule.completed"

	// New subscription schedule is created.
	EventSubscriptionScheduleCreated = "subscription_schedule.created"

	// 7 days before a subscription schedule will expire.
	EventSubscriptionScheduleExpiring = "subscription_schedule.expiring"

	// New subscription schedule is released.
	EventSubscriptionScheduleReleased = "subscription_schedule.released"

	// Subscription schedule is updated.
	EventSubscriptionScheduleUpdated = "subscription_schedule.updated"

	// New tax rate is created.
	EventTaxRateCreated = "tax_rate.created"

	// Tax rate is updated.
	EventTaxRateUpdated = "tax_rate.updated"

	// Top-up is canceled.
	EventTopupCanceled = "topup.canceled"

	// Top-up is created.
	EventTopupCreated = "topup.created"

	// Top-up fails.
	EventTopupFailed = "topup.failed"

	// Top-up is reversed.
	EventTopupReversed = "topup.reversed"

	// Top-up succeeds.
	EventTopupSucceeded = "topup.succeeded"

	// Transfer is created.
	EventTransferCreated = "transfer.created"

	// Transfer failed.
	EventTransferFailed = "transfer.failed"

	// After a transfer is paid. For Instant Payouts, the event will typically
	// be sent within 30 minutes.
	EventTransferPaid = "transfer.paid"

	// Transfer is reversed, including partial reversals.
	EventTransferReversed = "transfer.reversed"

	// Transfer's description or metadata is updated.
	EventTransferUpdated = "transfer.updated"
)
