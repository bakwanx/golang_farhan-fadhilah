

CREATE TABLE `user` (
  `id_user` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `birthdate` date DEFAULT current_timestamp(),
  `status` varchar(50) NOT NULL,
  `gender` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
);

ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`);

ALTER TABLE `user`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;
