const config = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    scene: {
        preload:preload,
        create:create,
        update:update,
        menu:menu
    },
    physics:{
        default: 'arcade',
        arcade: {
            gravity:{y:0},
            debug:false
        }
    },
    scale:{
        autoCenter: 1
    },
    parent: "game"
}
var game = new Phaser.Game(config)

function preload () {
    
    //this.game.scale.autoCenter = 1;
    this.load.image('snake',"/assets/images/snake.png");
    this.load.image('gapple',"/assets/images/apple.png");
    this.load.image('rapple',"/assets/images/redApple.png");
    this.load.image('vWall',"/assets/images/wall.png");
    this.load.image('neck',"/assets/images/snake.png"); 
    this.load.image('background',"/assets/images/background.jpg"); 
    this.load.image('play',"/assets/images/playLogo.png"); 
}
var apples = 0
var tail = [];
var dead = false;
var speed = 200;
function create () {
    
    //default world elements
    scoreText = this.add.text(16, 16, 'score: 0', { fontSize: '32px', fill: '#ffff' });
    cursors = this.input.keyboard.createCursorKeys();
    this.body = this.physics.add.group();
    this.neck = this.physics.add.group();
    // Snake creation
    snake = this.physics.add.sprite(400, 300, 'snake');
    snake.setCollideWorldBounds(true);
    snake.setScale(1.2)

    snake.checkWorldBounds = true;
    
    //variable elements
    Rapple = this.physics.add.sprite(Phaser.Math.Between(0, this.game.canvas.width-100),Phaser.Math.Between(0, this.game.canvas.height-100),'rapple') 
    Gapple = this.physics.add.sprite(Phaser.Math.Between(0, this.game.canvas.width-100),Phaser.Math.Between(0, this.game.canvas.height-100),'gapple')     

    //Interactions
    this.physics.add.collider(snake, this.body.getChildren(),die, null, this);
    this.physics.add.overlap(snake, Gapple, collectGreenApple, null, this);
    this.physics.add.overlap(snake, Rapple, collectRedApple, null, this); 
}

function gameOver (game){
    menu(game);
    game.scene.pause();
}
function die(){
    dead = true
}
function collectGreenApple (snake, gapple)
{   //
    //Apple reaction
    //
    //collectRedApple()
    gapple.disableBody(true, true);
    apples += 1;
    scoreText.setText('apples: ' + apples);

    Gapple = this.physics.add.sprite(Phaser.Math.Between(0, this.game.canvas.width+1),Phaser.Math.Between(0, this.game.canvas.height+1),'gapple')     
    this.physics.add.overlap(snake, Gapple, collectGreenApple, null, this);
    //
    //Add tail elements
    //
    tail = this.body.getChildren();
    if(this.neck.getChildren().length>(snake.width/speed)*100){ 
        this.body.create(
            this.neck.getChildren()[this.neck.getChildren().length-1].x, 
                this.neck.getChildren()[this.neck.getChildren().length-1].y, 'neck');                        
    }
    else{ 
        this.neck.create(snake.x, snake.y, 'neck');
    }
    //
    // 15% chance to generate a redApple when eating
    //
    odd = Phaser.Math.Between(0,100)
    if(odd<=15){
        Rapple = this.physics.add.sprite(Phaser.Math.Between(0, this.game.canvas.width-100),Phaser.Math.Between(0, this.game.canvas.height-100),'rapple') 
        this.physics.add.overlap(snake, Rapple, collectRedApple, null, this);
    }
}

function collectRedApple (snake,rapple)
{   
    rapple.disableBody(true, true);
    odd = Phaser.Math.Between(0,100)
    if(this.neck.getChildren().length<2){
        console.log("You died because you can't be smaller")
        die()
    }else{
        if(this.body.getChildren().length>2){
            for (let i =0; i<2 ; i++) {
                let removed = this.body.getChildren()[this.body.getChildren().length-1];
                removed.disableBody(true,true)
                this.body.remove(this.body.getChildren().length-1);
            }
        }else if(this.body.getChildren().length>0){
            let count = 0;
            for (let i = 0; i<this.body.getChildren().length-1 ; i++) {
                let removed = this.body.getChildren()[this.body.getChildren().length-1];
                removed.disableBody(true,true)
                this.body.remove(this.body.remove(this.body.getChildren().length-1));
                count++;
            }
            if(count!=2){
                let removed = this.neck.getChildren()[this.neck.getChildren().length-1];
                removed.disableBody(true,true)
                this.neck.remove(this.neck.getChildren()[this.neck.getChildren().length-1]);
            }
        }else{
            for (let i = 0; i<2 ; i++) {
                let removed = this.neck.getChildren()[this.neck.getChildren().length-1];
                removed.disableBody(true,true)
                this.neck.remove(this.neck.getChildren()[this.neck.getChildren().length-1]);
            }
        }
    }
    //
    // 15% chance to generate a redApple when eating
    //
    if(odd<=100){
        Rapple = this.physics.add.sprite(Phaser.Math.Between(0, this.game.canvas.width-100),Phaser.Math.Between(0, this.game.canvas.height-100),'rapple') 
        this.physics.add.overlap(snake, Rapple, collectRedApple, null, this);
    }  

}

function update(){  
    
    if (cursors.left.isDown & snake.body.velocity.x!=speed) {
        snake.body.setVelocityY(0);
        snake.body.setVelocityX(-speed);
    } else if (cursors.right.isDown & snake.body.velocity.x!=-speed) {
        snake.body.setVelocityY(0);
        snake.body.setVelocityX(speed);
    } else if (cursors.up.isDown & snake.body.velocity.y!=speed) {
        snake.body.setVelocityY(-speed);
        snake.body.setVelocityX(0);
    } else if(cursors.down.isDown & snake.body.velocity.y!=-speed){
        snake.body.setVelocityX(0);
        snake.body.setVelocityY(speed);
    } 
    //Create a "tail" efect
    if(this.neck.getChildren().length>0){      
        Phaser.Actions.ShiftPosition(this.neck.getChildren(), snake.x, snake.y,1);
    }
    if(tail.length>0){

        Phaser.Actions.ShiftPosition(tail, this.neck.getChildren()[this.neck.getChildren().length-1].x, this.neck.getChildren()[this.neck.getChildren().length-1].y);
    }
    //borders
    if(this.game.canvas.width-9==snake.x || snake.x ==9 || this.game.canvas.height-9==snake.y|| snake.y ==9 ){
        dead = true
    }
    if(dead){
        gameOver(this);
    }
}

function menu(game){
    var play = game.add.sprite(this.game.canvas.width/2,this.game.canvas.height/2,'play').setDepth(0);
    play.setScale(0.5)
}