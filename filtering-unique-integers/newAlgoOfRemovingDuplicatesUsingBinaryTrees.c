/*
  Copyright Junior ADI

  License GPLv3
*/
#include <stdio.h>
#include <stdlib.h>

// Définition de la structure d'un noeud d'arbre binaire
typedef struct Node {
    int data;
    struct Node* left;
    struct Node* right;
} Node;

// Fonction pour créer un nouveau noeud
Node* createNode(int data) {
    Node* newNode = (Node*)malloc(sizeof(Node));
    newNode->data = data;
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
    }
    
    return root;
}

// Fonction pour supprimer les doublons
Node* removeDuplicates(int* arr, int size) {
    Node* root = NULL;
    
    // Insérer les éléments dans l'arbre
    for (int i = 0; i < size; i++) {
        root = insertNode(root, arr[i]);
    }
    
    return root;
}

int main() {
    int arr[] = {1, 2, 3, 2, 4, 5, 1, 6};
    int size = sizeof(arr) / sizeof(arr[0]);
    
    Node* uniqueRoot = removeDuplicates(arr, size);
    
    // Parcourir l'arbre et afficher les éléments uniques
    printf("Eléments uniques:\n");
    if (uniqueRoot != NULL) {
        if (uniqueRoot->left != NULL) {
            printf("%d\n", uniqueRoot->left->data);
        }
        printf("%d\n", uniqueRoot->data);
        if (uniqueRoot->right != NULL) {
            printf("%d\n", uniqueRoot->right->data);
        }
    }
    
    return 0;
}
