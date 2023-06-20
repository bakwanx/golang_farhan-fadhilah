
CREATE TABLE `transaction` (
  `id_transaction` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_id` int(11) NOT NULL,
  `courier_id` int(11) NOT NULL
);

ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id_transaction`),
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD UNIQUE KEY `payment_id` (`payment_id`,`courier_id`),
  ADD KEY `courier_id` (`courier_id`);

ALTER TABLE `transaction`
  MODIFY `id_transaction` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `transaction`
  ADD CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id_user`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `transaction_ibfk_2` FOREIGN KEY (`courier_id`) REFERENCES `courier` (`id_courier`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `transaction_ibfk_3` FOREIGN KEY (`payment_id`) REFERENCES `payment` (`id_payment`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;
