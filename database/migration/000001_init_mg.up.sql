CREATE TABLE `users` (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `admin` tinyint(1) NOT NULL DEFAULT 0,
    `hash` varchar(60) NOT NULL,
    `requestAdmin` tinyint(1) NOT NULL DEFAULT 0
);

CREATE TABLE `books` (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `bookName` varchar(255) NOT NULL,
    `quantity` int NOT NULL DEFAULT 1,
    `availableQuantity` int NOT NULL DEFAULT 1
);

CREATE TABLE `requests` (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `status` ENUM('issued', 'request-issue', 'request-return'),
    `bookId` int NOT NULL,
    `userId` int NOT NULL,
    FOREIGN KEY (`bookId`) REFERENCES `books`(`id`),
    FOREIGN KEY (`userId`) REFERENCES `users`(`id`)
);