<template>
    <v-dialog lazy v-model="show" persistent max-width="350" class="p-account-delete-dialog" @keydown.esc="cancel">
        <v-card raised elevation="24">
            <v-container fluid class="pb-2 pr-2 pl-2">
                <v-layout row wrap>
                    <v-flex xs3 text-xs-center>
                        <v-icon size="54" color="grey lighten-1">delete_outline</v-icon>
                    </v-flex>
                    <v-flex xs9 text-xs-left align-self-center>
                        <div class="subheading pr-1"><translate>Are you sure you want to delete this account?</translate></div>
                    </v-flex>
                    <v-flex xs12 text-xs-right class="pt-3">
                        <v-btn @click.stop="cancel" depressed color="grey lighten-3" class="action-cancel">
                            <translate>Cancel</translate>
                        </v-btn>
                        <v-btn color="blue-grey lighten-2" depressed dark @click.stop="confirm"
                               class="action-confirm"><translate>Delete</translate>
                        </v-btn>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-card>
    </v-dialog>
</template>
<script>
    export default {
        name: 'p-account-delete-dialog',
        props: {
            show: Boolean,
            model: Object,
        },
        data() {
            return {
                loading: false,
            }
        },
        methods: {
            cancel() {
                this.$emit('cancel');
            },
            confirm() {
                this.loading = true;

                this.model.remove().then(() => {
                    this.loading = false;
                    this.$emit('confirm');
                });
            },
        }
    }
</script>
