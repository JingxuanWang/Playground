#include <iostream>
#include <stdio.h>
#include <string>

using namespace std;

const unsigned int PLAYER_SIDE = 1;
const unsigned int ENEMY_SIDE = 2;
const int MAX_TURN = 3;
const int CHARA_NUM = 3;

int player_chara_remain = CHARA_NUM;
int enemy_chara_remain = CHARA_NUM;

class chara {
	int side; // 1, player; 2, enemy
	int atk;
	int def;
	int hp;
	
	void hurt(int damage) {
		this->hp -= damage;
	}

	public:
	void attack(chara * enemy) {
		if (enemy == NULL || !this->isAlive() || !enemy->isAlive()) {
			return;
		}

		int damage = 1;	
		if (this->atk - enemy->def > 0) {
			damage = this->atk - enemy->def;
		}
		enemy->hurt(damage);
	}
	bool isAlive() {
		return this->hp > 0;
	}
	bool isPlayerChara() {
		return this->side == PLAYER_SIDE;
	}
	bool isEnemyChara() {
		return this->side == ENEMY_SIDE;
	}
};

void pre_turn() {}
void post_turn() {}
void pre_action() {}
void post_action() {}
void init_player_chara(chara *) {};
void init_enemy_chara(chara *) {};

int action(chara* src, chara* dst) {
	for (int a = 0; a < CHARA_NUM; ++a) {
		pre_action();
		// skip if already dead
		if (!src[a].isAlive()) {
			continue;
		}
		
		for (int b = 0; b < CHARA_NUM; ++b) {
			if (!dst[b].isAlive()) {
				continue;
			}

			// attack
			src[a].attack(&dst[b]);

			// if not dead
			if (dst[b].isAlive()) {
				continue;
			}

			if (dst[b].isPlayerChara()) {
				--player_chara_remain;

				// battle finished
				if (player_chara_remain <= 0) {
					return 2;
				}
			} else if (dst[b].isEnemyChara()) {
				--enemy_chara_remain;

				// battle finished
				if (enemy_chara_remain <= 0) {
					return 1;
				}
			} else {
				// unknown chara
			}
		
			break;
		}
		post_action();
	}
	return 0;
}

int battle_start() {

	// initialized player and enemy chara array
	chara * player_charas = new chara[CHARA_NUM];
	chara * enemy_charas = new chara[CHARA_NUM];

	init_player_chara(player_charas);
	init_enemy_chara(enemy_charas);

	int result = 0;

	for (int i = 0 ; i < MAX_TURN; ++i) {
		pre_turn();
		result = action(player_charas, enemy_charas);
		if (result) {
			return result;
		}

		result = action(enemy_charas, player_charas);
		if (result) {
			return result;
		}
		post_turn();
	}
	return result;
}


int main() {

	return 0;
}

