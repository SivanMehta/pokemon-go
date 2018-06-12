library(ggplot2)
library(dplyr)

data <- read.csv("data.csv")

data %>%
  ggplot() +
  aes(x = generation) +
  geom_line(aes(y = hp), color = 'red') +
  geom_line(aes(y = atk), color = 'blue') +
  geom_line(aes(y = def), color = 'blue') +
  geom_line(aes(y = spatk), color = 'green') +
  geom_line(aes(y = spdef), color = 'green') +
  geom_line(aes(y = spd), color = 'yellow') +
  
  labs(title = "Progression of Each Stat", x = "Generation", y = "Base")

ggsave("stat-progression.png", width = 10, height = 5)
