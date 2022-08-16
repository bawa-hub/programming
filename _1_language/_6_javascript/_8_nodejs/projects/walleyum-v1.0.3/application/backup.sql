INSERT INTO `currencies` (`id`, `name`, `symbol`, `icon`, `rateUsd`, `ratefromApi`, `crypto`, `metadata`, `active`, `createdAt`, `updatedAt`) VALUES
(1, 'US Dollar', 'USD', NULL, 1, 1, 0, NULL, 1, '2021-10-24 09:38:42', '2022-01-03 05:00:02');

INSERT INTO `gateways` (`id`, `name`, `value`, `apiKey`, `secretKey`, `email`, `isCrypto`, `active`, `isExchangePayment`, `ex1`, `ex2`, `createdAt`, `updatedAt`) VALUES
(1, 'Mollie', 'mollie', NULL, NULL, NULL, 0, 0, 0, NULL, NULL, '2021-10-24 09:38:42', '2021-10-23 06:08:49'),
(2, 'Coinbase', 'coinbase', NULL, NULL, NULL, 1, 0, 0, NULL, NULL, '2021-10-24 09:38:42', '2021-09-26 11:54:26'),
(3, 'Coin Payments', 'coinpayments', NULL, NULL, NULL, 1, 0, 0, NULL, NULL, '2021-10-24 09:38:42', '2021-09-26 16:37:39'),
(4, 'Paypal', 'paypal', NULL, NULL, NULL, 0, 0, 0, 'live', NULL, '2021-10-24 09:38:42', '2021-09-10 15:17:54'),
(5, 'Stripe', 'stripe', NULL, NULL, NULL, 0, 0, 0, NULL, NULL, '2021-10-24 09:38:42', '2021-09-24 19:09:34'),
(6, 'Coingate', 'coingate', NULL, NULL, NULL, 1, 0, 0, 'live', NULL, '2021-10-24 09:38:42', '2021-10-24 09:38:42'),
(7, 'Paystack', 'paystack', NULL, NULL, NULL, 0, 0, 0, NULL, NULL, '2021-10-24 09:38:42', '2021-10-24 09:38:42'),
(8, 'VoguePay', 'voguepay', NULL, NULL, NULL, 0, 0, 0, NULL, NULL, '2021-10-24 09:38:42', '2021-10-24 09:38:42');

INSERT INTO `pages` (`id`, `type`, `name`, `slug`, `content`, `createdAt`, `updatedAt`) VALUES
(1, 'landing', 'Home', 'home', '[]', '2021-10-24 09:38:42', '2022-01-02 10:55:46');

INSERT INTO `settings` (`id`, `value`, `param1`, `param2`, `createdAt`, `updatedAt`) VALUES
(1, 'refferal', '10', 'USD', '2021-10-24 09:38:42', '2021-10-24 09:38:42'),
(2, 'site', 'Walleyum', 'walleyum@tdevs.co', '2021-10-24 09:38:42', '2021-10-23 08:24:59'),
(3, 'appUrl', 'https://walleyum.tdevs.co', NULL, '2021-10-24 09:38:42', '2021-10-23 08:24:59'),
(4, 'apiUrl', 'https://walleyum.tdevs.co/api', NULL, '2021-10-24 09:38:42', '2021-10-23 08:24:59'),
(5, 'freecurrencyapi', NULL, NULL, '2021-10-24 09:38:42', '2021-10-24 09:38:42'),
(6, 'adjustments', '2.5', '0', '2021-10-24 09:38:42', '2021-10-24 09:56:50'),
(7, 'mainmenu', '[]', NULL, '2021-10-24 09:38:42', '2021-10-27 11:23:11'),
(8, 'footermenu', '[]', NULL, '2021-10-24 09:38:42', '2021-12-22 12:37:34'),
(9, 'tagline', 'Wallet, Exchanger, Cryptocurrency', NULL, '2021-10-24 09:38:42', '2021-10-24 09:38:42'),
(10, 'logo', NULL, NULL, '2021-10-24 09:38:42', '2021-10-27 09:25:44'),
(11, 'services', '[]', NULL, '2021-10-24 09:38:42', '2021-12-23 10:13:33'),
(12, 'solutions', '[]', NULL, '2021-10-24 09:38:42', '2021-12-22 17:10:46'),
(13, 'work', '[]', NULL, '2021-10-24 09:38:42', '2021-10-24 09:38:42'),
(14, 'faq', '[]', NULL, '2021-10-24 09:38:42', '2021-12-23 10:12:26');