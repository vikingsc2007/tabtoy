// Generated by github.com/vikingsc2007/tabtoy
// Version: 2.9.0
// DO NOT EDIT!!
#include <vector>
#include <map>
#include <string>

namespace table
{
	
	// Defined in table: Globals
	enum class ActorType
	{
		
		
		Leader = 0, // 唐僧
		
		
		Monkey = 1, // 孙悟空
		
		
		Pig = 2, // 猪八戒
		
		
		Hammer = 3, // 沙僧
	
	};
	

	

	
	

	

	// Defined in table: Globals
	class Vec2
	{
	public:
	
		
		public:
 		int X_ = 0; 
	
		
		public:
 		int Y_ = 0; 
	

	}; 
	

	

	// Defined in table: Sample
	class Prop
	{
	public:
	
		
		public:
 		int HP_ = 10; // 血量
	
		
		public:
 		float AttackRate_ = 0.0f; // 攻击速率
	
		
		public:
 		ActorType ExType_ = ActorType::Leader; // 额外类型
	

	}; 
	

	

	// Defined in table: Sample
	class AttackParam
	{
	public:
	
		
		public:
 		int Value_ = 0; // 攻击值
	

	}; 
	

	

	// Defined in table: Sample
	class SampleDefine
	{
	public:
	
		/// <summary> 
		/// 唯一ID
		/// </summary>
		public:
 		long long ID_ = 0; 
	
		/// <summary> 
		/// 名称
		/// </summary>
		public:
 		std::string Name_ = ""; 
	
		/// <summary> 
		/// 图标ID
		/// </summary>
		public:
 		int IconID_ = 0; 
	
		/// <summary> 
		/// 攻击率
		/// </summary>
		public:
 		float NumericalRate_ = 0.0f; 
	
		/// <summary> 
		/// 物品id
		/// </summary>
		public:
 		int ItemID_ = 100; 
	
		/// <summary> 
		/// BuffID
		/// </summary>
		public:
 		std::vector<int> BuffID_; 
	
		/// <summary> 
		/// 位置
		/// </summary>
		public:
 		Vec2 Pos_; 
	
		/// <summary> 
		/// 类型
		/// </summary>
		public:
 		ActorType Type_ = ActorType::Leader; 
	
		/// <summary> 
		/// 技能ID列表
		/// </summary>
		public:
 		std::vector<int> SkillID_; 
	
		/// <summary> 
		/// 攻击参数
		/// </summary>
		public:
 		AttackParam AttackParam_; 
	
		/// <summary> 
		/// 单结构解析
		/// </summary>
		public:
 		Prop SingleStruct_; 
	
		/// <summary> 
		/// 字符串结构
		/// </summary>
		public:
 		std::vector<Prop> StrStruct_; 
	

	}; 
	

	
	

	// Defined in table: Config	
	class Config
	{
	
	public:
		tabtoy::Logger TableLogger;
	
		
		/// <summary> 
		/// Sample
		/// </summary>
		public:
 		std::vector<SampleDefine> Sample_; 
	
	
		//#region Index code
	 	std::map<long long, SampleDefine> _SampleByID;
	public:
		class SampleDefine* GetSampleByID(long long ID, SampleDefine* def = nullptr)
        {
            auto ret = _SampleByID.find( ID );
            if ( ret != _SampleByID.end() )
            {
                return &ret->second;
            }
			
			if ( def == nullptr )
			{
				TableLogger.ErrorLine("GetSampleByID failed, ID: %s", ID);
			}

            return def;
        }
		std::map<std::string, SampleDefine> _SampleByName;
	public:
		class SampleDefine* GetSampleByName(std::string Name, SampleDefine* def = nullptr)
        {
            auto ret = _SampleByName.find( Name );
            if ( ret != _SampleByName.end() )
            {
                return &ret->second;
            }
			
			if ( def == nullptr )
			{
				TableLogger.ErrorLine("GetSampleByName failed, Name: %s", Name);
			}

            return def;
        }
		
	
		//#endregion
		//#region Deserialize code
		
	public:
		static void Deserialize( Config& ins, tabtoy::DataReader& reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0xa0000:
                	{
						ins.Sample_.emplace_back( reader.ReadStruct<SampleDefine>(Deserialize) );
                	}
                	break; 
                }
             }

			
			// Build Sample Index
			for( size_t i = 0;i< ins.Sample_.size();i++)
			{
				auto element = ins.Sample_[i];
				
				ins._SampleByID.emplace(std::make_pair(element.ID_, element));
				
				ins._SampleByName.emplace(std::make_pair(element.Name_, element));
				
			}
			
		}
	public:
		static void Deserialize( Vec2& ins, tabtoy::DataReader& reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0x10000:
                	{
						ins.X_ = reader.ReadInt32();
                	}
                	break; 
                	case 0x10001:
                	{
						ins.Y_ = reader.ReadInt32();
                	}
                	break; 
                }
             }

			
		}
	public:
		static void Deserialize( Prop& ins, tabtoy::DataReader& reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0x10000:
                	{
						ins.HP_ = reader.ReadInt32();
                	}
                	break; 
                	case 0x50001:
                	{
						ins.AttackRate_ = reader.ReadFloat();
                	}
                	break; 
                	case 0x80002:
                	{
						ins.ExType_ = reader.ReadEnum<ActorType>();
                	}
                	break; 
                }
             }

			
		}
	public:
		static void Deserialize( AttackParam& ins, tabtoy::DataReader& reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0x10000:
                	{
						ins.Value_ = reader.ReadInt32();
                	}
                	break; 
                }
             }

			
		}
	public:
		static void Deserialize( SampleDefine& ins, tabtoy::DataReader& reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0x20000:
                	{
						ins.ID_ = reader.ReadInt64();
                	}
                	break; 
                	case 0x60001:
                	{
						ins.Name_ = reader.ReadString();
                	}
                	break; 
                	case 0x10002:
                	{
						ins.IconID_ = reader.ReadInt32();
                	}
                	break; 
                	case 0x50003:
                	{
						ins.NumericalRate_ = reader.ReadFloat();
                	}
                	break; 
                	case 0x10004:
                	{
						ins.ItemID_ = reader.ReadInt32();
                	}
                	break; 
                	case 0x10005:
                	{
						ins.BuffID_.emplace_back( reader.ReadInt32() );
                	}
                	break; 
                	case 0x90006:
                	{
						ins.Pos_ = reader.ReadStruct<Vec2>(Deserialize);
                	}
                	break; 
                	case 0x80007:
                	{
						ins.Type_ = reader.ReadEnum<ActorType>();
                	}
                	break; 
                	case 0x10008:
                	{
						ins.SkillID_.emplace_back( reader.ReadInt32() );
                	}
                	break; 
                	case 0x90009:
                	{
						ins.AttackParam_ = reader.ReadStruct<AttackParam>(Deserialize);
                	}
                	break; 
                	case 0x9000a:
                	{
						ins.SingleStruct_ = reader.ReadStruct<Prop>(Deserialize);
                	}
                	break; 
                	case 0x9000b:
                	{
						ins.StrStruct_.emplace_back( reader.ReadStruct<Prop>(Deserialize) );
                	}
                	break; 
                }
             }

			
		}
		//#endregion
	

	};
	
	
	
	
	
	
	
	
	
	
}
