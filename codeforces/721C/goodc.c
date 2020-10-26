#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

struct edge {
	int v, t;
	struct edge *next;
};

struct edge **vv;

void add(int u, int v, int t) {
	struct edge *x = malloc(sizeof(*x));

	x->v = v;
	x->t = t;
	x->next = vv[u];
	vv[u] = x;
}

int pp[5001][5000], tt[5001][5000];

void print(int k, int u) {
	if (k > 0) {
		print(k - 1, pp[k][u]);
		printf("%d ", u + 1);
	}
}

int main() {
	int k, u, v, t, tmax, n, m;

	scanf("%d%d%d", &n, &m, &tmax);
	vv = calloc(n, sizeof(*vv));
	while (m-- > 0) {
		scanf("%d%d%d", &u, &v, &t);
		u--, v--;
		add(u, v, t);
	}
	for (k = 1; k <= n; k++)
		for (u = 0; u < n; u++)
			tt[k][u] = tmax + 1;
	tt[1][0] = 0;
	for (k = 1; k <= n - 1; k++){
		for (u = 0; u < n; u++){
			if (tt[k][u] <= tmax) {
				struct edge *e;

				for (e = vv[u]; e != NULL; e = e->next) {
					int v = e->v, t = e->t;

					if (tt[k][u] + t <= tmax){
						if (tt[k + 1][v] > tt[k][u] + t) {
							tt[k + 1][v] = tt[k][u] + t;
							pp[k + 1][v] = u;
						}
					}
				}
			}
		}
	}
	for (k = n; k >= 1; k--)
		if (tt[k][n - 1] != tmax + 1)
			break;
	printf("%d\n", k);
	print(k, n - 1);
	printf("\n");
	return 0;
}