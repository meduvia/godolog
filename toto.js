const { CommandInteraction, MessageEmbed } = require("discord.js");

module.exports = {
    name: "clear",
    description: "Supprime les messages.",
    permission: "MANAGE_MESSAGES",
    options: [
        {
            name: "nombre",
            description: "Choisis le nombre de messages à supprimer.",
            type: "NUMBER",
            require: true
        },
        {
            name: "cible",
            description: "Choisis une cible à qui il faut supprimer les messages.",
            type: "USER",
            require: false
        },
    ],
    /**
     * 
     * @param {CommandInteraction} interaction 
     */
    async execute(interaction) {
        const { channel, options } = interaction;

        const Amount = options.getNumber("nombre");
        const Target = options.getMember("cible");

        const Messages = await channel.messages.fetch();

        const Response = new MessageEmbed()
        .setColor("BLUE");

        if(Target) {
            let i = 0;
            const filtered = [];
            (await Messages). filter((m) => {
                if(m.author.id === Target.id && Amount > i) {
                    filtered.push(m);
                    i++;
                }
            })

            await channel.bulkDelete(filtered, true).then(messages => {
                Response.setDescription(`${messages.size} messages de ${Target} ont été supprimés`)
                interaction.reply({embeds: [Response]});
            })
        } else {
            await channel.bulkDelete(Amount, true).then(messages => {
                Response.setDescription(`${messages.size} messages du salon ont été supprimés`)
                interaction.reply({embeds: [Response]});
            }
                )
        }
        
    }
}