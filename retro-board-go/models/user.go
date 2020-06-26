package models

import (
	"github.com/jinzhu/gorm"
)

// @Entity({ name: 'users' })
// @Index(['username', 'accountType'], { unique: true })
// export default class User {
//   @PrimaryColumn({ primary: true, generated: false, unique: true })
//   public id: string;
//   @Column()
//   @Index()
//   public name: string;
//   @Column({ default: 'anonymous' })
//   public accountType: AccountType;
//   @Column({ nullable: true, type: 'character varying' })
//   public username: string | null;
//   @Column({ nullable: true, type: 'character varying' })
//   public photo: string | null;
//   @Column({ nullable: false, type: 'character varying', default: 'en' })
//   public language: string;
//   @ManyToOne(() => SessionTemplate, { nullable: true, eager: false })
//   public defaultTemplate: SessionTemplate | null | undefined;
//   @CreateDateColumn({ type: 'timestamp with time zone' })
//   public created: Date | undefined;
//   @UpdateDateColumn({ type: 'timestamp with time zone' })
//   public updated: Date | undefined;
//   constructor(id: string, name: string) {
//     this.id = id;
//     this.name = name;
//     this.language = 'en';
//     this.accountType = 'anonymous';
//     this.username = null;
//     this.photo = null;
//   }
// }

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	AccountType string `gorm:"default:'anonymous'"`
	Username    string
	Photo       string
	Language    string `gorm:"not null"`

	// Birthday     *time.Time
	// Email        string  `gorm:"type:varchar(100);unique_index"`
	// Role         string  `gorm:"size:255"` // set field size to 255
	// MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
	// Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	// Address      string  `gorm:"index:addr"` // create index with name `addr` for address
	// IgnoreMe     int     `gorm:"-"` // ignore this field
}
