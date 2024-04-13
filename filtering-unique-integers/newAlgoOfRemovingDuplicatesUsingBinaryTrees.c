/*
  Copyright Junior ADI

  License GPLv3
*/

#include <stdio.h>
#include <stdlib.h>

// Définition de la structure d'un noeud d'arbre binaire
typedef struct Node {
    int data;
    int count;
    struct Node* left;
    struct Node* right;
} Node;

// Fonction pour créer un nouveau noeud
Node* createNode(int data) {
    Node* newNode = (Node*)malloc(sizeof(Node));
    newNode->data = data;
    newNode->count = 1;
    newNode->left = NULL;
    newNode->right = NULL;
    return newNode;
}

// Fonction pour insérer un élément dans l'arbre
Node* insertNode(Node* root, int data) {
    if (root == NULL) {
        return createNode(data);
    }
    
    if (data < root->data) {
        root->left = insertNode(root->left, data);
    } else if (data > root->data) {
        root->right = insertNode(root->right, data);
    } else {
        root->count++;
    }
    
    return root;
}

// Fonction pour afficher les éléments non uniques
void printNonUnique(Node* root) {
    if (root != NULL) {
        printNonUnique(root->left);
        if (root->count > 1) {
            printf("%d apparaît %d fois\n", root->data, root->count);
        }
        printNonUnique(root->right);
    }
}

// Fonction pour afficher les éléments uniques
void printUnique(Node* root) {
    if (root != NULL) {
        printUnique(root->left);
        if (root->count == 1) {
            printf("%d est unique\n", root->data);
        }
        printUnique(root->right);
    }
}

int main() {
    int arr[] = {1, 2, 3, 2, 4, 5, 1, 6};
    int size = sizeof(arr) / sizeof(arr[0]);
    
    Node* root = NULL;
    
    // Insérer les éléments dans l'arbre
    for (int i = 0; i < size; i++) {
        root = insertNode(root, arr[i]);
    }
    
    // Parcourir l'arbre et afficher les éléments non uniques
    printf("Eléments non uniques:\n");
    printNonUnique(root);

    // Parcourir l'arbre et afficher les éléments non uniques
    printf("Eléments uniques:\n");
    printUnique(root);
    
    return 0;
}
